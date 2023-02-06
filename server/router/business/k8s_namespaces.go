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

func RegisterK8SNamespacesRouter(router *gin.RouterGroup) {
	nsRouter := router.Group("/ns")

	k8sNamespaces := businessv1.NewK8SNameSpaces()
	{
		nsRouter.GET("/:id", k8sNamespaces.Get)
		nsRouter.GET("", k8sNamespaces.GetAll)
		nsRouter.POST("", k8sNamespaces.Add)
		nsRouter.PUT("/:id", k8sNamespaces.Update)
		nsRouter.PUT("", k8sNamespaces.UpdateAll)
		nsRouter.DELETE("/:id", k8sNamespaces.DeleteOne)
		nsRouter.DELETE("", k8sNamespaces.DeleteMore)
	}
}
