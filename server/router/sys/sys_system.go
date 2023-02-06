package sys

import (
	"github.com/gin-gonic/gin"
	sysv1 "github.com/liuzw3018/kubei2dou/server/api/v1/sys"
)

/**
 * @Author: liu zw
 * @Date: 2022/11/10 13:38
 * @File: sys_account.go
 * @Description: //TODO $
 * @Version:
 */

func RegisterSysSystemRouter(router *gin.RouterGroup) {
	systemRouter := router.Group("system")

	sysSystem := sysv1.NewSysAccount()
	sysAccountRouter := systemRouter.Group("account")
	sysDeptRouter := systemRouter.Group("dept")
	{
		sysAccountRouter.GET("list", sysSystem.AccountList)
	}
	{
		sysDeptRouter.GET("list", sysSystem.DeptList)
	}
}
