package sys

import (
	"github.com/gin-gonic/gin"
	sysv1 "github.com/liuzw3018/kubei2dou/server/api/v1/sys"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/20 17:49
 * @File: sys_admin.go
 * @Description: //TODO $
 * @Version:
 */

func RegisterSysAdminRouter(router *gin.RouterGroup) {
	adminRouter := router.Group("/admin")

	sysAdmin := sysv1.NewSysAdmin()
	{
		adminRouter.GET("/:id", sysAdmin.Get)
		adminRouter.GET("", sysAdmin.GetAll)
		adminRouter.POST("", sysAdmin.Add)
		adminRouter.PUT("/:id", sysAdmin.Update)
		adminRouter.PUT("", sysAdmin.UpdateAll)
		adminRouter.DELETE("/:id", sysAdmin.DeleteOne)
		adminRouter.DELETE("", sysAdmin.DeleteMore)
	}
}
