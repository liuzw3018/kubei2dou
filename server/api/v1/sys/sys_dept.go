package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/kubei2dou/server/api"
)

/**
 * @Author: liu zw
 * @Date: 2022/11/10 14:59
 * @File: sys_dept.go
 * @Description: //TODO $
 * @Version:
 */

type sysDept struct{}

func NewSysDept() api.BaseApiHandle {
	return &sysDept{}
}

func (sd *sysDept) GetAll(c *gin.Context) {

}

func (sd *sysDept) Get(c *gin.Context) {

}

func (sd *sysDept) Update(c *gin.Context) {

}

func (sd *sysDept) UpdateAll(c *gin.Context) {

}

func (sd *sysDept) Add(c *gin.Context) {

}

func (sd *sysDept) DeleteOne(c *gin.Context) {

}

func (sd *sysDept) DeleteMore(c *gin.Context) {

}
