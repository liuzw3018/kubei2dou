package sys

import (
	"github.com/gin-gonic/gin"
	sysv1 "github.com/liuzw3018/kubei2dou/server/api/v1/sys"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/25 18:21
 * @File: sys_login.go
 * @Description: //TODO $
 * @Version:
 */

func RegisterSysLoginRouter(router *gin.RouterGroup) {
	loginRouter := router.Group("/")

	sysLogin := sysv1.NewSysLogin()
	{
		loginRouter.POST("login", sysLogin.Login)
		loginRouter.GET("logout", sysLogin.Logout)
	}
}
