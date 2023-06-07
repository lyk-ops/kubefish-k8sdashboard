package k8s

import "github.com/pddzl/kubefish/server/service"

type ApiGroup struct {
	NodeApi
	ResourceApi
	NamespaceApi
	PodApi
	ReplicaSetApi
	DeploymentApi
	DaemonSetApi
}

var (
	nodeService       = service.ServiceGroupApp.K8sServiceGroup.NodeService
	resourceService   = service.ServiceGroupApp.K8sServiceGroup.ResourceService
	namespaceService  = service.ServiceGroupApp.K8sServiceGroup.NamespaceService
	podService        = service.ServiceGroupApp.K8sServiceGroup.PodService
	replicaSetService = service.ServiceGroupApp.K8sServiceGroup.ReplicaSetService
	deploymentService = service.ServiceGroupApp.K8sServiceGroup.DeploymentService
	daemonSetService  = service.ServiceGroupApp.K8sServiceGroup.DaemonSetService
)
