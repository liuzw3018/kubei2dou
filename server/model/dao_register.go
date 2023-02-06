package model

import (
	"fmt"
	"github.com/liuzw3018/kubei2dou/server/model/dao"
	"github.com/liuzw3018/kubei2dou/server/pkg/utils"
	"github.com/liuzw3018/saber/lib"
	uuid "github.com/satori/go.uuid"
)

/**
 * @Author: liuzw
 * @Date: 2023/2/2 10:56
 * @File: dao_register.go
 * @Description: //TODO $
 * @Version:
 */

func RegisterDao() {

	if err := lib.GORMDefaultPool.AutoMigrate(dao.Admin{}); err != nil {
		panic(err)
	}
	initAdminUser()
}

func initAdminUser() {
	admin := &dao.Admin{
		Username:    "admin",
		RealName:    "管理员",
		NickName:    "管理员",
		Salt:        "admin",
		Avatar:      "https://q1.qlogo.cn/g?b=qq&nk=190848757&s=640",
		HomePath:    "/dashboard/analysis",
		UUID:        uuid.NewV4(),
		Password:    utils.GenSaltPassword("admin", "123456"),
		Description: "admin",
		IsDelete:    0,
	}
	fmt.Println("注册管理员用户")
	lib.GORMDefaultPool.Create(admin)
	fmt.Println("注册管理员用户完成")
}
