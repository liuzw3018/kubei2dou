package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/kubei2dou/server/model/dto"
	"github.com/liuzw3018/kubei2dou/server/pkg/public"
	"github.com/liuzw3018/kubei2dou/server/pkg/response"
	"net/http"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/25 18:20
 * @File: sys_user.go
 * @Description: //TODO $
 * @Version:
 */

type sysUser struct{}

func NewSysUser() *sysUser {
	return &sysUser{}
}

// Login godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /api/v1/login
// @Accept json
// @Produce json
// @Param body body dto.SysUserLoginInput true "body"
// @Success 200 {object} response.Response{data=dto.SysUserLoginOutput} "success"
// @Router /api/v1/login [post]
func (su *sysUser) Login(c *gin.Context) {
	var (
		code  public.ResponseCode
		token string
		out   = &dto.SysUserLoginOutput{}

		err error
	)
	params := &dto.SysUserLoginInput{}
	if err = params.BindingValidParams(c); err != nil {
		code = 2001
		goto ERR
	}

	if !(params.Username == "admin" && params.Password == "123456") {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}

	token, err = public.GenToken(public.Users{Username: "admin", Password: "123456"})
	if err != nil {
		code = 2002
		goto ERR
	}

	out.Token = token
	out.Role = "admin"

	response.Success(c, public.SuccessCode, "ok", out)
	return

ERR:
	response.Error(c, code, err)
	return
}

func (su *sysUser) Logout(c *gin.Context) {

}
