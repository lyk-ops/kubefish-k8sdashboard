package cluster

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubefish/server/api"
)

type NamespaceRouter struct{}

func (nr *NamespaceRouter) InitNamespaceRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	namespaceRouter := Router.Group("namespace")
	namespaceApi := api.ApiGroupApp.K8sApiGroup.ClusterApi.NamespaceApi
	{
		namespaceRouter.POST("getNamespaces", namespaceApi.GetNamespaces)
		namespaceRouter.GET("getNamespaceDetail", namespaceApi.GetNamespaceDetail)
		namespaceRouter.DELETE("deleteNamespace", namespaceApi.DeleteNamespace)
		namespaceRouter.GET("getNamespaceName", namespaceApi.GetNamespaceName)
		namespaceRouter.POST("updateNamespace", namespaceApi.UpdateNamespace)
	}
	return namespaceRouter
}
