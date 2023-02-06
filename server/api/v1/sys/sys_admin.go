package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/kubei2dou/server/api"
	"github.com/liuzw3018/kubei2dou/server/model/dao"
	"github.com/liuzw3018/kubei2dou/server/model/dto"
	"github.com/liuzw3018/kubei2dou/server/pkg/public"
	"github.com/liuzw3018/kubei2dou/server/pkg/response"
	"github.com/liuzw3018/saber/lib"
	"log"
	"strconv"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/20 17:40
 * @File: sys_admin.go
 * @Description: //TODO $
 * @Version:
 */

type sysAdmin struct{}

func NewSysAdmin() api.BaseApiHandle {
	return &sysAdmin{}
}

// GetAll godoc
// @Summary 管理员配置
// @Description 管理员配置
// @Tags 管理员接口
// @UserID /api/v1/admin
// @Accept json
// @Produce json
// @Param body body dto.SysAdminListInput true "body"
// @Success 200 {object} response.Response{data=dto.SysAdminListOutput} "success"
// @Router /api/v1/admin [get]
func (sa *sysAdmin) GetAll(c *gin.Context) {
	params := &dto.SysAdminListInput{}
	if err := params.BindingValidParams(c); err != nil {
		response.Error(c, public.InvalidRequestErrorCode, err)
		return
	}
	adminInfo := &dao.Admin{}
	list, total, err := adminInfo.PageList(c, lib.GORMDefaultPool, params)
	if err != nil {
		response.Error(c, public.OtherErrorCode, err)
		return
	}

	var outList []dto.SysAdminListOnceInfoOutput
	for _, listItem := range list {
		var onceOut dto.SysAdminListOnceInfoOutput
		onceOut.Nickname = listItem.NickName
		onceOut.Account = listItem.Username
		onceOut.CreateTime = listItem.CreatedAt
		onceOut.Email = listItem.Email
		onceOut.ID = listItem.Id
		onceOut.Remark = listItem.Description
		onceOut.Status = listItem.IsDelete
		outList = append(outList, onceOut)
	}

	out := &dto.SysAdminListOutput{
		Total: total,
		Items: outList,
	}

	response.Success(c, public.SuccessCode, "OK", out)
}

func (sa *sysAdmin) Get(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		response.Error(c, public.InternalErrorCode, err)
		return
	}
	admin := &dao.Admin{Id: userId}
	admin, err = admin.Find(c, lib.GORMDefaultPool, admin)
	if err != nil {
		log.Println(err)
	}

	out := &dto.SysAdminInfoOutput{
		Username: admin.Username,
		Avatar:   admin.Avatar,
		UUID:     admin.UUID,
		RealName: admin.RealName,
		NickName: admin.NickName,
		HomePath: admin.HomePath,
		Token:    c.Request.Header.Get("Authorization"),
		Role: dto.SysAdminRoleOutput{
			RoleName: "Super Admin",
			Value:    "super",
		},
		Description: admin.Description,
	}
	response.Success(c, public.SuccessCode, "OK", out)
}

func (sa *sysAdmin) Update(c *gin.Context) {

}

func (sa *sysAdmin) UpdateAll(c *gin.Context) {

}

func (sa *sysAdmin) Add(c *gin.Context) {

}

func (sa *sysAdmin) DeleteOne(c *gin.Context) {

}

func (sa *sysAdmin) DeleteMore(c *gin.Context) {

}
