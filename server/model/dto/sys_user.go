package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/kubei2dou/server/pkg/public"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/26 14:49
 * @File: sys_user.go
 * @Description: //TODO $
 * @Version:
 */

type SysUserLoginInput struct {
	Username string `json:"username" form:"username" comment:"用户名" example:"admin" validate:"valNotNull"`
	Password string `json:"password"  form:"password" comment:"密码" example:"123" validate:"valNotNull"`
}

type SysUserLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"" validate:""`
	Role  string `json:"role" form:"role" comment:"角色组" example:"" validate:""`
}

func (params *SysUserLoginInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}
