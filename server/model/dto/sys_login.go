package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/kubei2dou/server/pkg/public"
)

/**
 * @Author: liuzw
 * @Date: 2023/2/2 11:00
 * @File: sys_login.go
 * @Description: //TODO $
 * @Version:
 */

type SysAdminLoginInput struct {
	Username string `json:"username" form:"username" comment:"用户名" example:"admin" validate:"valNotNull"`
	Password string `json:"password"  form:"password" comment:"密码" example:"123" validate:"valNotNull"`
}

func (params *SysAdminLoginInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type SysAdminLoginOutput struct {
	UserId int                `json:"userId" form:"userId" comment:"用户ID" validate:""`
	Token  string             `json:"token" form:"token" comment:"token" validate:""`
	Role   SysAdminRoleOutput `json:"role" form:"role" comment:"角色组" validate:""`
}
