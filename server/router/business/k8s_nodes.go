package business

import (
	"github.com/gin-gonic/gin"
	businessv1 "github.com/liuzw3018/kubei2dou/server/api/v1/business"
)

/**
 * @Author: liuzw
 * @Date: 2023/2/3 09:52
 * @File: k8s_namespaces.go
 * @Description: //TODO $
 * @Version:
 */

func RegisterK8SNodesRouter(router *gin.RouterGroup) {
	nodesRouter := router.Group("/nodes")

	k8sNodes := businessv1.NewK8SNodes()
	{
		nodesRouter.GET("/:id", k8sNodes.Get)
		nodesRouter.GET("", k8sNodes.GetAll)
		nodesRouter.POST("", k8sNodes.Add)
		nodesRouter.PUT("/:id", k8sNodes.Update)
		nodesRouter.PUT("", k8sNodes.UpdateAll)
		nodesRouter.DELETE("/:id", k8sNodes.DeleteOne)
		nodesRouter.DELETE("", k8sNodes.DeleteMore)
	}
}
