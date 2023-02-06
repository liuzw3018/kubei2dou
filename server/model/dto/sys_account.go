package dto

/**
 * @Author: liuzw
 * @Date: 2023/2/2 11:00
 * @File: sys_account.go
 * @Description: //TODO $
 * @Version:
 */

type SysSystemAccountListInput struct {
	Page     int `json:"page" form:"page" comment:"页数" example:"" validate:""`
	PageSize int `json:"pageSize" form:"pageSize" comment:"每页条数" example:"" validate:""`
}

type SysSystemAccountListOutput struct {
	Total int                      `json:"total" form:"total" comment:"总条数" example:"" validate:""`
	Items []SysSystemAccountOutput `json:"items" form:"items" comment:"账号数据" validate:""`
}

type SysSystemAccountOutput struct {
	ID         string `json:"id" form:"id" comment:"用户ID" example:"" validate:""`
	Account    string `json:"account" form:"account" comment:"用户名" example:"" validate:""`
	Email      string `json:"email" form:"email" comment:"用户邮箱" example:"" validate:""`
	Nickname   string `json:"nickname" form:"nickname" comment:"用户昵称" example:"" validate:""`
	Role       int    `json:"role" form:"role" comment:"用户角色" example:"" validate:""`
	CreateTime string `json:"createTime" form:"createTime" comment:"用户创建时间" example:"" validate:""`
	Remark     string `json:"remark" form:"remark" comment:"备注" example:"" validate:""`
	Status     int    `json:"status" form:"status" comment:"用户状态" example:"" validate:""`
}
