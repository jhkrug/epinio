package application

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/epinio/epinio/helpers/kubernetes"
	"github.com/epinio/epinio/internal/services"
	"github.com/epinio/epinio/pkg/api/core/v1/models"

	pkgerrors "github.com/pkg/errors"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/util/retry"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

type AppServiceBind struct {
	service  string // name of the service getting bound
	resource string // name of the kube secret to mount as volume to make the service params available in the app
}

type AppServiceBindList []AppServiceBind

// Workload manages applications that are deployed. It provides workload
// (deployments) specific actions for the application model.
type Workload struct {
	app     models.AppRef
	cluster *kubernetes.Cluster
}

// NewWorkload constructs and returns a workload representation from an application reference.
func NewWorkload(cluster *kubernetes.Cluster, app models.AppRef) *Workload {
	return &Workload{cluster: cluster, app: app}
}

func ToBinds(ctx context.Context, services services.ServiceList, appName string, userName string) (AppServiceBindList, error) {
	bindings := AppServiceBindList{}

	for _, service := range services {
		bindResource, err := service.GetBinding(ctx, appName, userName)
		if err != nil {
			return AppServiceBindList{}, err
		}
		bindings = append(bindings, AppServiceBind{
			resource: bindResource.Name,
			service:  service.Name(),
		})
	}

	return bindings, nil
}

func (b AppServiceBindList) ToVolumesArray() []corev1.Volume {
	volumes := []corev1.Volume{}

	for _, binding := range b {
		volumes = append(volumes, corev1.Volume{
			Name: binding.service,
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: binding.resource,
				},
			},
		})
	}

	return volumes
}

func (b AppServiceBindList) ToMountsArray() []corev1.VolumeMount {
	mounts := []corev1.VolumeMount{}

	for _, binding := range b {
		mounts = append(mounts, corev1.VolumeMount{
			Name:      binding.service,
			ReadOnly:  true,
			MountPath: fmt.Sprintf("/services/%s", binding.service),
		})
	}

	return mounts
}

// BoundServicesChange imports the currently bound services into the deployment. It takes a ServiceList, not just
// names, as it has to create/retrieve the associated service binding secrets. It further takes a set of the old
// services. This enables incremental modification of the deployment (add, remove affected, instead of wholsesale
// replacement).
func (a *Workload) BoundServicesChange(ctx context.Context, userName string, oldServices NameSet, newServices services.ServiceList) error {
	_, err := Get(ctx, a.cluster, a.app)
	if err != nil {
		// Should not happen. Application was validated to exist
		// already somewhere by callers.
		return err
	}

	bindings, err := ToBinds(ctx, newServices, a.app.Name, userName)
	if err != nil {
		return err
	}

	// Create name-keyed maps from old/new slices for quick lookup and decision. No linear searches.

	new := map[string]struct{}{}

	for _, s := range newServices {
		new[s.Name()] = struct{}{}
	}

	// Read, modify and write the deployment
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		deployment, err := a.Deployment(ctx)
		if err != nil {
			return err
		}

		// The action is done in multiple iterations over the deployment's volumes and volumemounts.
		// The first iteration over each determines removed services (in old, not in new). The second
		// iteration, over the new services now, adds all which are not in old, i.e. actually new.

		newVolumes := []corev1.Volume{}
		newMounts := []corev1.VolumeMount{}

		for _, volume := range deployment.Spec.Template.Spec.Volumes {
			_, hasold := oldServices[volume.Name]
			_, hasnew := new[volume.Name]

			// Note that volumes which are not in old are passed and kept. These are the volumes
			// not related to services.

			if hasold && !hasnew {
				continue
			}

			newVolumes = append(newVolumes, volume)
		}

		// TODO: Iterate over containers and find the one matching the app name
		for _, mount := range deployment.Spec.Template.Spec.Containers[0].VolumeMounts {

			_, hasold := oldServices[mount.Name]
			_, hasnew := new[mount.Name]

			// Note that volumes which are in not in old are passed and kept. These are the volumes
			// not related to services.

			if hasold && !hasnew {
				continue
			}

			newMounts = append(newMounts, mount)
		}

		for _, binding := range bindings {
			// Skip services which already exist
			if _, hasold := oldServices[binding.service]; hasold {
				continue
			}

			newVolumes = append(newVolumes, corev1.Volume{
				Name: binding.service,
				VolumeSource: corev1.VolumeSource{
					Secret: &corev1.SecretVolumeSource{
						SecretName: binding.resource,
					},
				},
			})

			newMounts = append(newMounts, corev1.VolumeMount{
				Name:      binding.service,
				ReadOnly:  true,
				MountPath: fmt.Sprintf("/services/%s", binding.service),
			})
		}

		// Write the changed set of mounts and volumes back to the deployment ...
		deployment.Spec.Template.Spec.Volumes = newVolumes
		deployment.Spec.Template.Spec.Containers[0].VolumeMounts = newMounts

		// ... and then the cluster.
		_, err = a.cluster.Kubectl.AppsV1().Deployments(a.app.Org).Update(
			ctx, deployment, metav1.UpdateOptions{})

		return err
	})
}

// EnvironmentChange imports the current environment into the
// deployment. This requires only the names of the currently existing
// environment variables, not the values, as the import is internally
// done as pod env specifications using secret key references.
func (a *Workload) EnvironmentChange(ctx context.Context, varNames []string) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		deployment, err := a.Deployment(ctx)
		if err != nil {
			return err
		}

		evSecretName := a.app.MakeEnvSecretName()

		// 1. Remove all the old EVs referencing the app's EV secret.
		// 2. Add entries for the new set of EV's (S.a varNames).
		// 3. Replace container spec
		//
		// Note: While 1+2 could be optimized to only remove entries of
		//       EVs not in varNames, and add only entries for varNames
		//       not in Env, this is way more complex for what is likely
		//       just 10 entries. I expect any gain in perf to be
		//       negligible, and completely offset by the complexity of
		//       understanding and maintaining it later. Full removal
		//       and re-adding is much simpler to understand, and should
		//       be fast enough.

		newEnvironment := []corev1.EnvVar{}

		for _, ev := range deployment.Spec.Template.Spec.Containers[0].Env {
			// Drop EV if pulled from EV secret of the app
			if ev.ValueFrom != nil &&
				ev.ValueFrom.SecretKeyRef != nil &&
				ev.ValueFrom.SecretKeyRef.Name == evSecretName {
				continue
			}
			// Keep everything else.
			newEnvironment = append(newEnvironment, ev)
		}

		for _, varName := range varNames {
			newEnvironment = append(newEnvironment, corev1.EnvVar{
				Name: varName,
				ValueFrom: &corev1.EnvVarSource{
					SecretKeyRef: &corev1.SecretKeySelector{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: evSecretName,
						},
						Key: varName,
					},
				},
			})
		}

		deployment.Spec.Template.Spec.Containers[0].Env = newEnvironment

		_, err = a.cluster.Kubectl.AppsV1().Deployments(a.app.Org).Update(
			ctx, deployment, metav1.UpdateOptions{})

		return err
	})
}

// Scale changes the number of instances (replicas) for the
// application's Deployment.
func (a *Workload) Scale(ctx context.Context, instances int32) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		deployment, err := a.Deployment(ctx)
		if err != nil {
			return err
		}

		deployment.Spec.Replicas = &instances

		_, err = a.cluster.Kubectl.AppsV1().Deployments(a.app.Org).Update(
			ctx, deployment, metav1.UpdateOptions{})

		return err
	})
}

// Deployment is a helper, it returns the kube deployment resource of the workload.
func (a *Workload) Deployment(ctx context.Context) (*appsv1.Deployment, error) {
	return a.cluster.Kubectl.AppsV1().Deployments(a.app.Org).Get(
		ctx, a.app.Name, metav1.GetOptions{},
	)
}

// Metrics returns CPU and Memory usage or and error if one occurs
func (a *Workload) Metrics(ctx context.Context) (int64, int64, error) {
	mc, err := metrics.NewForConfig(a.cluster.RestConfig)
	if err != nil {
		return 0, 0, err
	}

	deployment, err := a.Deployment(ctx)
	if err != nil {
		return 0, 0, err
	}

	selector := labels.Set(deployment.Spec.Selector.MatchLabels).AsSelector().String()
	podMetrics, err := mc.MetricsV1beta1().PodMetricses(a.app.Org).
		List(ctx, metav1.ListOptions{LabelSelector: selector})
	if err != nil {
		return 0, 0, err
	}

	cpuUsage := resource.NewQuantity(0, resource.DecimalSI)
	memUsage := resource.NewQuantity(0, resource.BinarySI)

	for _, podMetric := range podMetrics.Items {
		podContainers := podMetric.Containers
		for _, container := range podContainers {
			// TODO: Filter only application containers?
			cpuUsage.Add(*container.Usage.Cpu())
			memUsage.Add(*container.Usage.Memory())
		}
	}

	// cpu * 1000 -> milliCPUs (rounded)
	milliCPUs := int64(math.Round(cpuUsage.ToDec().AsApproximateFloat64() * 1000))

	mem, ok := memUsage.AsInt64()
	if !ok {
		fmt.Printf("memUsage.AsDec = %T %+v\n", memUsage.AsDec(), memUsage.AsDec())
	}

	return milliCPUs, mem, nil
}

// Restarts returns the number of restarts the application currently has
// The pod's restarts is the sum of restarts of each and every container in that Pod
func (a *Workload) Restarts(ctx context.Context) (int32, error) {
	deployment, err := a.Deployment(ctx)
	if err != nil {
		return 0, err
	}

	selector := labels.Set(deployment.Spec.Selector.MatchLabels).AsSelector().String()
	pods, err := a.cluster.Kubectl.CoreV1().Pods(a.app.Org).List(ctx,
		metav1.ListOptions{LabelSelector: selector})
	if err != nil {
		return 0, err
	}

	restarts := int32(0)
	for _, p := range pods.Items {
		for _, cs := range p.Status.ContainerStatuses {
			if cs.Name == a.app.Name {
				restarts += cs.RestartCount
			}
		}
	}

	return restarts, nil
}

// Get returns the state of the app deployment encoded in the workload.
func (a *Workload) Get(ctx context.Context, deployment *appsv1.Deployment) *models.AppDeployment {
	active := false
	routes := []string{}
	stageID := ""
	status := ""
	username := ""

	// Query application deployment for stageID and status (ready vs desired replicas)

	deploymentSelector := fmt.Sprintf("app.kubernetes.io/part-of=%s,app.kubernetes.io/name=%s", a.app.Org, a.app.Name)

	deploymentListOptions := metav1.ListOptions{
		LabelSelector: deploymentSelector,
	}

	deployments, err := a.cluster.Kubectl.AppsV1().Deployments(a.app.Org).List(ctx, deploymentListOptions)

	desiredReplicas := int32(0)
	readyReplicas := int32(0)

	var createdAt time.Time

	if err != nil {
		status = pkgerrors.Wrap(err, "failed to get Deployment status").Error()
	} else if len(deployments.Items) < 1 {
		status = "0/0"
	} else {
		desiredReplicas = deployments.Items[0].Status.Replicas
		readyReplicas = deployments.Items[0].Status.ReadyReplicas

		createdAt = deployments.Items[0].ObjectMeta.CreationTimestamp.Time

		status = fmt.Sprintf("%d/%d",
			deployments.Items[0].Status.ReadyReplicas,
			deployments.Items[0].Status.Replicas)

		stageID = deployments.Items[0].
			Spec.Template.ObjectMeta.Labels["epinio.suse.org/stage-id"]
		username = deployments.Items[0].Spec.Template.ObjectMeta.Labels["app.kubernetes.io/created-by"]

		active = true
	}

	routes, err = ListRoutes(ctx, a.cluster, a.app)
	if err != nil {
		routes = []string{err.Error()}
	}

	restarts, err := a.Restarts(ctx)
	if err != nil {
		status = pkgerrors.Wrap(err, "failed to calculate application restarts").Error()
	}

	cpu, mem, err := a.Metrics(ctx)
	if err != nil {
		status = pkgerrors.Wrap(err, "failed to calculate resource usage").Error()
	}

	return &models.AppDeployment{
		Active:          active,
		CreatedAt:       createdAt.Format(time.RFC3339), // ISO 8601
		MilliCPUs:       cpu,
		MemoryBytes:     mem,
		Restarts:        restarts,
		Username:        username,
		StageID:         stageID,
		Status:          status,
		Routes:          routes,
		DesiredReplicas: desiredReplicas,
		ReadyReplicas:   readyReplicas,
	}
}
