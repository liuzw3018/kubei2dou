package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/kubei2dou/server/pkg/public"
	uuid "github.com/satori/go.uuid"
	"time"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/26 14:03
 * @File: sys_admin.go
 * @Description: //TODO $
 * @Version:
 */

type SysAdminListInput struct {
	Page  int    `json:"page" form:"page" comment:"页数" example:"1" validate:"valNotNull"`
	Limit int    `json:"pageSize" form:"pageSize" comment:"每页条数" example:"10" validate:"valNotNull"`
	Info  string `json:"info" form:"info" comment:"关键字" example:"" validate:""`
}

func (params *SysAdminListInput) BindingValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type SysAdminInfoOutput struct {
	Username    string             `json:"username" form:"username" comment:"用户登录名" example:"" validate:""`
	Avatar      string             `json:"avatar" form:"avatar" comment:"用户头像" example:"" validate:""`
	UUID        uuid.UUID          `json:"userId" form:"userId" comment:"用户ID" example:"" validate:""`
	HomePath    string             `json:"homePath" form:"homePath" comment:"用户主页" example:"" validate:""`
	RealName    string             `json:"realName" form:"realName" comment:"用户全名" example:"" validate:""`
	NickName    string             `json:"nickName" form:"nickName" comment:"用户昵称" example:"" validate:""`
	Token       string             `json:"token" form:"token" comment:"token" example:"" validate:""`
	Role        SysAdminRoleOutput `json:"role" form:"role" comment:"角色组" example:"" validate:""`
	Description string             `json:"desc" form:"desc" comment:"用户描述" example:"" validate:""`
}

type SysAdminRoleOutput struct {
	RoleName string `json:"roleName" form:"roleName" comment:"角色名" example:"" validate:""`
	Value    string `json:"value"  form:"value" comment:"角色" example:"" validate:""`
}

type SysAdminListOutput struct {
	Total int64                        `json:"total" form:"total" comment:"总条数"  validate:""`
	Items []SysAdminListOnceInfoOutput `json:"items" form:"items" comment:"账号数据"  validate:""`
}

type SysAdminListOnceInfoOutput struct {
	ID         int       `json:"id" form:"id" comment:"用户ID"  validate:""`
	Account    string    `json:"account" form:"account" comment:"用户名"  validate:""`
	Email      string    `json:"email" form:"email" comment:"用户邮箱"  validate:""`
	Nickname   string    `json:"nickname" form:"nickname" comment:"用户昵称"  validate:""`
	Role       int       `json:"role" form:"role" comment:"用户角色"  validate:""`
	CreateTime time.Time `json:"createTime" form:"createTime" comment:"用户创建时间"  validate:""`
	Remark     string    `json:"remark" form:"remark" comment:"备注"  validate:""`
	Status     int       `json:"status" form:"status" comment:"用户状态"  validate:""`
}
