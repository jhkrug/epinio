// Package services encapsulates all the functionality around Epinio services
package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/epinio/epinio/helpers/kubernetes"
	epinioerrors "github.com/epinio/epinio/internal/errors"
	"github.com/epinio/epinio/internal/organizations"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ServiceList []*Service

// Service contains the information needed for Epinio to address a specific service.
type Service struct {
	SecretName string
	OrgName    string
	Service    string
	Username   string
	kubeClient *kubernetes.Cluster
}

// Lookup locates a Service by namespace and name. It finds the Service
// instance by looking for the relevant Secret.
func Lookup(ctx context.Context, kubeClient *kubernetes.Cluster, namespace, service string) (*Service, error) {
	// TODO 844 inline

	secretName := serviceResourceName(namespace, service)

	s, err := kubeClient.GetSecret(ctx, namespace, secretName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, errors.New("service not found")
		}
		return nil, err
	}
	username := s.ObjectMeta.Labels["app.kubernetes.io/created-by"]

	return &Service{
		SecretName: secretName,
		OrgName:    namespace,
		Service:    service,
		kubeClient: kubeClient,
		Username:   username,
	}, nil
}

// List returns a ServiceList of all available Services in the specified namespace. If no namespace is
// specified (empty string) then services across all namespaces are returned.
func List(ctx context.Context, cluster *kubernetes.Cluster, namespace string) (ServiceList, error) {
	labelSelector := "app.kubernetes.io/name=epinio"

	// Verify namespace, if specified
	if namespace != "" {
		exists, err := organizations.Exists(ctx, cluster, namespace)
		if err != nil {
			return ServiceList{}, err
		}
		if !exists {
			return ServiceList{}, epinioerrors.NamespaceMissingError{
				Namespace: namespace,
			}
		}

		labelSelector = fmt.Sprintf("%s, epinio.suse.org/namespace=%s", labelSelector, namespace)
	}

	secrets, err := cluster.Kubectl.CoreV1().
		Secrets(namespace).List(ctx,
		metav1.ListOptions{
			LabelSelector: labelSelector,
		})

	if err != nil {
		return nil, err
	}

	result := ServiceList{}

	for _, s := range secrets.Items {
		service := s.ObjectMeta.Labels["epinio.suse.org/service"]
		namespace := s.ObjectMeta.Labels["epinio.suse.org/namespace"]
		username := s.ObjectMeta.Labels["app.kubernetes.io/created-by"]

		secretName := s.ObjectMeta.Name

		result = append(result, &Service{
			SecretName: secretName,
			OrgName:    namespace,
			Service:    service,
			kubeClient: cluster,
			Username:   username,
		})
	}

	return result, nil
}

// CreateService creates a new  service instance from namespace,
// name, and a map of parameters.
func CreateService(ctx context.Context, cluster *kubernetes.Cluster, name, namespace, username string,
	data map[string]string) (*Service, error) {

	secretName := serviceResourceName(namespace, name)

	_, err := cluster.GetSecret(ctx, namespace, secretName)
	if err == nil {
		return nil, errors.New("Service of this name already exists.")
	}

	// Convert from `string -> string` to the `string -> []byte` expected
	// by kube.
	sdata := make(map[string][]byte)
	for k, v := range data {
		sdata[k] = []byte(v)
	}

	err = cluster.CreateLabeledSecret(ctx, namespace, secretName, sdata,
		map[string]string{
			// "epinio.suse.org/service-type": "custom",
			"epinio.suse.org/service":      name,
			"epinio.suse.org/namespace":    namespace,
			"app.kubernetes.io/name":       "epinio",
			"app.kubernetes.io/created-by": username,
			// "app.kubernetes.io/version":     cmd.Version
			// FIXME: Importing cmd causes cycle
			// FIXME: Move version info to separate package!
		},
	)
	if err != nil {
		return nil, err
	}
	return &Service{
		SecretName: secretName,
		OrgName:    namespace,
		Service:    name,
		kubeClient: cluster,
	}, nil
}

// Name returns the service instance's name
func (s *Service) Name() string {
	return s.Service
}

// User returns the service's username
func (s *Service) User() string {
	return s.Username
}

// Org returns the service instance's namespace
func (s *Service) Org() string {
	return s.OrgName
}

// GetBinding returns the secret representing the instance's binding
// to the application. This is actually the instance's secret itself,
// independent of the application.
func (s *Service) GetBinding(ctx context.Context, appName string, _ string) (*corev1.Secret, error) {
	cluster := s.kubeClient
	serviceSecret, err := cluster.GetSecret(ctx, s.OrgName, s.SecretName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, errors.New("service does not exist")
		}
		return nil, err
	}

	return serviceSecret, nil
}

// Delete destroys the service instance, i.e. its underlying secret
// holding the instance's parameters
func (s *Service) Delete(ctx context.Context) error {
	return s.kubeClient.DeleteSecret(ctx, s.OrgName, s.SecretName)
}

// Details returns the service instance's configuration.
// I.e. the parameter data.
func (s *Service) Details(ctx context.Context) (map[string]string, error) {
	serviceSecret, err := s.kubeClient.GetSecret(ctx, s.OrgName, s.SecretName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, errors.New("service does not exist")
		}
		return nil, err
	}

	details := map[string]string{}

	for k, v := range serviceSecret.Data {
		details[k] = string(v)
	}

	return details, nil
}

// serviceResourceName returns a name for a kube service resource
// representing the namespace and service
func serviceResourceName(namespace, service string) string {
	return fmt.Sprintf("service.org-%s.svc-%s", namespace, service)
}
