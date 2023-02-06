package sys

import (
	"github.com/gin-gonic/gin"
	sysv1 "github.com/liuzw3018/kubei2dou/server/api/v1/sys"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/25 18:21
 * @File: sys_user.go
 * @Description: //TODO $
 * @Version:
 */

func RegisterSysUserRouter(router *gin.RouterGroup) {
	baseRouter := router.Group("/")

	sysBase := sysv1.NewSysUser()
	{
		baseRouter.POST("login", sysBase.Login)
		baseRouter.POST("logout", sysBase.Logout)
	}
}
