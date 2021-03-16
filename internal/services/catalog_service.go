package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/suse/carrier/deployments"
	"github.com/suse/carrier/internal/duration"
	"github.com/suse/carrier/internal/interfaces"
	"github.com/suse/carrier/kubernetes"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/dynamic"
)

// CatalogService is a Service created using Service Catalog.
// Implements the Service interface.
type CatalogService struct {
	InstanceName string
	OrgName      string
	Service      string
	Class        string
	Plan         string
	kubeClient   *kubernetes.Cluster
}

// CatalogServiceList returns a ServiceList of all available catalog Services
func CatalogServiceList(kubeClient *kubernetes.Cluster, org string) (interfaces.ServiceList, error) {
	labelSelector := fmt.Sprintf("app.kubernetes.io/name=carrier, carrier.suse.org/organization=%s", org)

	serviceInstanceGVR := schema.GroupVersionResource{
		Group:    "servicecatalog.k8s.io",
		Version:  "v1beta1",
		Resource: "serviceinstances",
	}

	dynamicClient, err := dynamic.NewForConfig(kubeClient.RestConfig)
	if err != nil {
		return nil, err
	}

	serviceInstances, err := dynamicClient.Resource(serviceInstanceGVR).
		Namespace(deployments.WorkloadsDeploymentID).
		List(context.Background(),
			metav1.ListOptions{
				LabelSelector: labelSelector,
			})

	if err != nil {
		return nil, err
	}

	result := interfaces.ServiceList{}

	for _, serviceInstance := range serviceInstances.Items {
		spec := serviceInstance.Object["spec"].(map[string]interface{})
		className := spec["clusterServiceClassExternalName"].(string)
		planName := spec["clusterServicePlanExternalName"].(string)

		metadata := serviceInstance.Object["metadata"].(map[string]interface{})
		instanceName := metadata["name"].(string)
		labels := metadata["labels"].(map[string]interface{})
		org := labels["carrier.suse.org/organization"].(string)
		service := labels["carrier.suse.org/service"].(string)

		result = append(result, &CatalogService{
			InstanceName: instanceName,
			OrgName:      org,
			Service:      service,
			Class:        className,
			Plan:         planName,
			kubeClient:   kubeClient,
		})
	}

	return result, nil
}

func CatalogServiceLookup(kubeClient *kubernetes.Cluster, org, service string) (interfaces.Service, error) {
	instanceName := serviceResourceName(org, service)

	serviceInstanceGVR := schema.GroupVersionResource{
		Group:    "servicecatalog.k8s.io",
		Version:  "v1beta1",
		Resource: "serviceinstances",
	}

	dynamicClient, err := dynamic.NewForConfig(kubeClient.RestConfig)
	if err != nil {
		return nil, err
	}

	serviceInstance, err := dynamicClient.Resource(serviceInstanceGVR).Namespace(deployments.WorkloadsDeploymentID).
		Get(context.Background(), instanceName, metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	spec := serviceInstance.Object["spec"].(map[string]interface{})
	className := spec["clusterServiceClassExternalName"].(string)
	planName := spec["clusterServicePlanExternalName"].(string)

	return &CatalogService{
		InstanceName: instanceName,
		OrgName:      org,
		Service:      service,
		Class:        className,
		Plan:         planName,
		kubeClient:   kubeClient,
	}, nil
}

func CreateCatalogService(kubeClient *kubernetes.Cluster, name, org, class, plan string, parameters map[string]string) (interfaces.Service, error) {
	resourceName := serviceResourceName(org, name)

	param, err := json.Marshal(parameters)
	if err != nil {
		return nil, err
	}

	data := fmt.Sprintf(`{
		"apiVersion": "servicecatalog.k8s.io/v1beta1",
		"kind": "ServiceInstance",
		"metadata": {
			"name": "%s",
			"namespace": "%s",
			"labels": {
				"carrier.suse.org/service-type": "catalog",
				"carrier.suse.org/service":      "%s",
				"carrier.suse.org/organization": "%s",
				"app.kubernetes.io/name":        "carrier"
			}
		},
		"spec": {
			"clusterServiceClassExternalName": "%s",
			"clusterServicePlanExternalName": "%s" },
		"parameters": %s
	}`, resourceName, deployments.WorkloadsDeploymentID,
		name, org, class, plan, param)

	decoderUnstructured := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	obj := &unstructured.Unstructured{}
	_, _, err = decoderUnstructured.Decode([]byte(data), nil, obj)
	if err != nil {
		return nil, err
	}

	serviceInstanceGVR := schema.GroupVersionResource{
		Group:    "servicecatalog.k8s.io",
		Version:  "v1beta1",
		Resource: "serviceinstances",
	}

	dynamicClient, err := dynamic.NewForConfig(kubeClient.RestConfig)
	if err != nil {
		return nil, err
	}

	// todo validations - check service instance existence

	_, err = dynamicClient.Resource(serviceInstanceGVR).Namespace(deployments.WorkloadsDeploymentID).
		Create(context.Background(),
			obj,
			metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return &CatalogService{
		InstanceName: resourceName,
		OrgName:      org,
		Service:      name,
		Class:        class,
		Plan:         plan,
		kubeClient:   kubeClient,
	}, nil
}

func (s *CatalogService) Name() string {
	return s.Service
}

func (s *CatalogService) Org() string {
	return s.OrgName
}

// GetBinding returns an application-specific secret for the service to be
// bound to that application.
func (s *CatalogService) GetBinding(appName string) (*corev1.Secret, error) {
	// TODO Label the secret

	bindingName := bindingResourceName(s.OrgName, s.Service, appName)

	binding, err := s.LookupBinding(bindingName)
	if err != nil {
		return nil, err
	}
	if binding == nil {
		_, err = s.CreateBinding(bindingName, s.OrgName, s.Service, appName)
		if err != nil {
			return nil, err
		}
	}

	return s.GetBindingSecret(bindingName)
}

// LookupBinding finds a ServiceBinding object for the application with Name
// appName if there is one.
func (s *CatalogService) LookupBinding(bindingName string) (interface{}, error) {
	serviceBindingGVR := schema.GroupVersionResource{
		Group:    "servicecatalog.k8s.io",
		Version:  "v1beta1",
		Resource: "servicebindings",
	}

	dynamicClient, err := dynamic.NewForConfig(s.kubeClient.RestConfig)
	if err != nil {
		return nil, err
	}

	serviceBinding, err := dynamicClient.Resource(serviceBindingGVR).Namespace(deployments.WorkloadsDeploymentID).
		Get(context.Background(), bindingName, metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return serviceBinding, nil
}

// CreateBinding creates a ServiceBinding for the application with name appName.
func (s *CatalogService) CreateBinding(bindingName, org, serviceName, appName string) (interface{}, error) {
	serviceInstanceName := serviceResourceName(org, serviceName)

	data := fmt.Sprintf(`{
		"apiVersion": "servicecatalog.k8s.io/v1beta1",
		"kind": "ServiceBinding",
		"metadata": { 
			"name": "%s", 
			"namespace": "%s",
		    "labels": { 
				"app.kubernetes.io/name": "%s",
				"app.kubernetes.io/part-of": "%s",
				"app.kubernetes.io/component": "servicebinding",
				"app.kubernetes.io/managed-by": "carrier"
			}
		},
		"spec": {
			"instanceRef": { "name": "%s" },
			"secretName": "%s" 
		}
	}`, bindingName, deployments.WorkloadsDeploymentID, appName, org, serviceInstanceName, bindingName)

	decoderUnstructured := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	obj := &unstructured.Unstructured{}
	_, _, err := decoderUnstructured.Decode([]byte(data), nil, obj)
	if err != nil {
		return nil, err
	}

	serviceBindingGVR := schema.GroupVersionResource{
		Group:    "servicecatalog.k8s.io",
		Version:  "v1beta1",
		Resource: "servicebindings",
	}

	dynamicClient, err := dynamic.NewForConfig(s.kubeClient.RestConfig)
	if err != nil {
		return nil, err
	}

	serviceBinding, err := dynamicClient.Resource(serviceBindingGVR).Namespace(deployments.WorkloadsDeploymentID).
		Create(context.Background(), obj, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	// Update the binding secret with kubernetes app labels
	secret, err := s.GetBindingSecret(bindingName)
	if err != nil {
		return nil, err
	}

	labels := secret.GetLabels()
	if labels == nil {
		labels = map[string]string{}
	}
	labels["app.kubernetes.io/name"] = appName
	labels["app.kubernetes.io/part-of"] = org
	labels["app.kubernetes.io/component"] = "servicebindingsecret"
	labels["app.kubernetes.io/managed-by"] = "carrier"
	secret.SetLabels(labels)

	_, err = s.kubeClient.Kubectl.CoreV1().Secrets(deployments.WorkloadsDeploymentID).Update(context.Background(), secret, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	return serviceBinding, nil
}

// GetBindingSecret returns the Secret that represents the binding of a Service
// to an Application.
func (s *CatalogService) GetBindingSecret(bindingName string) (*corev1.Secret, error) {
	return s.kubeClient.WaitForSecret(deployments.WorkloadsDeploymentID, bindingName,
		duration.ToServiceSecret())
}

// DeleteBinding deletes the ServiceBinding resource. The relevant secret will
// also be deleted automatically.
func (s *CatalogService) DeleteBinding(appName string) error {
	bindingName := bindingResourceName(s.OrgName, s.Service, appName)

	serviceBindingGVR := schema.GroupVersionResource{
		Group:    "servicecatalog.k8s.io",
		Version:  "v1beta1",
		Resource: "servicebindings",
	}

	dynamicClient, err := dynamic.NewForConfig(s.kubeClient.RestConfig)
	if err != nil {
		return err
	}

	return dynamicClient.Resource(serviceBindingGVR).Namespace(deployments.WorkloadsDeploymentID).
		Delete(context.Background(), bindingName, metav1.DeleteOptions{})
}

func (s *CatalogService) Delete() error {
	serviceInstanceGVR := schema.GroupVersionResource{
		Group:    "servicecatalog.k8s.io",
		Version:  "v1beta1",
		Resource: "serviceinstances",
	}

	dynamicClient, err := dynamic.NewForConfig(s.kubeClient.RestConfig)
	if err != nil {
		return err
	}

	return dynamicClient.Resource(serviceInstanceGVR).Namespace(deployments.WorkloadsDeploymentID).
		Delete(context.Background(), s.InstanceName, metav1.DeleteOptions{})
}

func (s *CatalogService) Status() (string, error) {
	serviceInstanceGVR := schema.GroupVersionResource{
		Group:    "servicecatalog.k8s.io",
		Version:  "v1beta1",
		Resource: "serviceinstances",
	}

	dynamicClient, err := dynamic.NewForConfig(s.kubeClient.RestConfig)
	if err != nil {
		return "", err
	}

	serviceInstance, err := dynamicClient.Resource(serviceInstanceGVR).Namespace(deployments.WorkloadsDeploymentID).
		Get(context.Background(), s.InstanceName, metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return "Not Found", nil
		} else {
			return "", err
		}
	}

	status := serviceInstance.Object["status"].(map[string]interface{})
	provisioned := status["provisionStatus"].(string)

	return provisioned, nil
}

func (s *CatalogService) WaitForProvision() error {
	serviceInstanceGVR := schema.GroupVersionResource{
		Group:    "servicecatalog.k8s.io",
		Version:  "v1beta1",
		Resource: "serviceinstances",
	}

	dynamicClient, err := dynamic.NewForConfig(s.kubeClient.RestConfig)
	if err != nil {
		return err
	}

	namespace := dynamicClient.Resource(serviceInstanceGVR).Namespace(deployments.WorkloadsDeploymentID)

	return wait.PollImmediate(time.Second, duration.ToServiceProvision(), func() (bool, error) {
		serviceInstance, err := namespace.Get(context.Background(), s.InstanceName, metav1.GetOptions{})
		if err != nil {
			if apierrors.IsNotFound(err) {
				return false, errors.New("Not Found")
			}
			return false, err
		}

		status, ok := serviceInstance.Object["status"].(map[string]interface{})
		if !ok {
			return false, nil
		}

		provisioned, ok := status["provisionStatus"].(string)
		if !ok {
			return false, nil
		}

		return provisioned == "Provisioned", nil
	})
}

func (s *CatalogService) Details() (map[string]string, error) {
	details := map[string]string{}

	details["Class"] = s.Class
	details["Plan"] = s.Plan

	return details, nil
}