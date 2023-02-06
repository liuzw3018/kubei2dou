package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/kubei2dou/server/model/dao"
	"github.com/liuzw3018/kubei2dou/server/model/dto"
	"github.com/liuzw3018/kubei2dou/server/pkg/public"
	"github.com/liuzw3018/kubei2dou/server/pkg/response"
	"github.com/liuzw3018/saber/lib"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/25 18:20
 * @File: sys_login.go
 * @Description: //TODO $
 * @Version:
 */

type sysLogin struct{}

func NewSysLogin() *sysLogin {
	return &sysLogin{}
}

// Login godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @UserID /api/v1/login
// @Accept json
// @Produce json
// @Param body body dto.SysAdminLoginInput true "body"
// @Success 200 {object} response.Response{data=dto.SysAdminLoginOutput} "success"
// @Router /api/v1/login [post]
func (su *sysLogin) Login(c *gin.Context) {
	var (
		code  public.ResponseCode
		token string
		out   *dto.SysAdminLoginOutput
		admin = &dao.Admin{}

		err error
	)
	params := &dto.SysAdminLoginInput{}
	if err = params.BindingValidParams(c); err != nil {
		code = 2001
		goto ERR
	}

	admin, err = admin.LoginCheck(c, lib.GORMDefaultPool, params)
	if err != nil {
		code = 2002
		goto ERR
	}

	token, err = public.GenToken(public.Users{
		Username: params.Username,
		Password: params.Password,
	})
	if err != nil {
		code = 2002
		goto ERR
	}

	out = &dto.SysAdminLoginOutput{
		UserId: admin.Id,
		Token:  token,
		Role: dto.SysAdminRoleOutput{
			RoleName: "Super Admin",
			Value:    "admin",
		},
	}

	response.Success(c, public.SuccessCode, "ok", out)
	return

ERR:
	response.Error(c, code, err)
	return
}

// Logout godoc
// @Summary 管理员退出
// @Description 管理员退出
// @Tags 管理员接口
// @UserID /api/v1/logout
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=dto.SysAdminLoginOutput} "success"
// @Router /api/v1/logout [get]
func (su *sysLogin) Logout(c *gin.Context) {

	response.Success(c, public.SuccessCode, "ok", nil)
}
