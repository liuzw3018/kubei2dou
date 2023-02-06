package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/kubei2dou/server/model/dto"
	"github.com/liuzw3018/kubei2dou/server/pkg/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/25 10:50
 * @File: sys_admin.go
 * @Description: //TODO $
 * @Version:
 */

type Admin struct {
	Id          int       `json:"id" gorm:"primary_key;type:int" description:"自增主键"`
	Username    string    `json:"Username" gorm:"column:user_name;type:varchar(100)" description:"管理员用户名"`
	RealName    string    `json:"RealName" gorm:"column:real_name;type:varchar(100)" description:"用户全名"`
	NickName    string    `json:"NickName" gorm:"column:nick_name;type:varchar(100)" description:"用户昵称"`
	Avatar      string    `json:"avatar" gorm:"column:avatar;type:varchar(100)" description:"用户头像"`
	HomePath    string    `json:"homePath" gorm:"column:home_path;type:varchar(100)" description:"用户主页"`
	Email       string    `json:"email" gorm:"column:email;type:varchar(100)" description:"用户邮箱"`
	Salt        string    `json:"Salt" gorm:"column:salt;type:varchar(100)" description:"盐"`
	UUID        uuid.UUID `json:"uuid" gorm:"column:uuid;unique" description:"用户uuid"`
	Role        int       `json:"role" gorm:"column:role" description:"用户角色"`
	DeptID      string    `json:"deptID" gorm:"column:dept_id" description:"用户角色"`
	Password    string    `json:"Password" gorm:"column:password;type:varchar(100)" description:"密码"`
	Description string    `json:"desc" gorm:"column:description" description:"用户描述"`
	UpdatedAt   time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt   time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete    int       `json:"is_delete" gorm:"column:is_delete;type:int" description:"是否删除"`
}

func (a *Admin) TableName() string {
	return "task_admin"
}

func (a *Admin) Find(c *gin.Context, tx *gorm.DB, search *Admin) (*Admin, error) {
	out := &Admin{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (a *Admin) LoginCheck(c *gin.Context, tx *gorm.DB, param *dto.SysAdminLoginInput) (*Admin, error) {
	adminInfo, err := a.Find(c, tx,
		&Admin{Username: param.Username, IsDelete: 0})
	if err != nil {
		return nil, errors.New("用户信息不存在")
	}
	saltPassword := utils.GenSaltPassword(adminInfo.Salt, param.Password)
	if adminInfo.Password != saltPassword {
		return nil, errors.New("密码错误，请重新输入")
	}
	return adminInfo, nil
}

func (a *Admin) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(a).Error
}

func (a *Admin) PageList(c *gin.Context, tx *gorm.DB, param *dto.SysAdminListInput) ([]Admin, int64, error) {
	var (
		total  int64
		list   []Admin
		offset = (param.Page - 1) * param.Limit
	)
	query := tx.WithContext(c)
	query = query.Table(a.TableName()).Where("is_delete=0")
	//if param.Info != "" {
	//	query = query.Where("service_name like ? or service_desc like ?", "%"+param.Info+"%", "%"+param.Info+"%")
	//}
	if err := query.Limit(param.Limit).
		Offset(offset).
		Order("id desc").
		Find(&list).
		Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	query.Limit(param.Limit).Offset(offset).Count(&total)
	return list, total, nil
}
