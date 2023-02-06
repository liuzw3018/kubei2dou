package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/kubei2dou/server/pkg/public"
	"github.com/liuzw3018/kubei2dou/server/pkg/response"
)

/**
 * @Author: liu zw
 * @Date: 2022/11/3 18:48
 * @File: sys_account.go
 * @Description: //TODO $ 文件可以删除
 * @Version:
 */

type sysAccount struct{}

func NewSysAccount() *sysAccount {
	return &sysAccount{}
}

func (ss *sysAccount) AccountList(c *gin.Context) {
	response.Success(c, public.SuccessCode, "OK", []interface{}{})
}

func (ss *sysAccount) DeptList(c *gin.Context) {
	response.Success(c, public.SuccessCode, "OK", []interface{}{})
}
