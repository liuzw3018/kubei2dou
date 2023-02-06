package api

import "github.com/gin-gonic/gin"

/**
 * @Author: liu zw
 * @Date: 2022/10/20 17:39
 * @File: BaseApi.go
 * @Description: //TODO $
 * @Version:
 */

type BaseApiHandle interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	UpdateAll(c *gin.Context)
	Add(c *gin.Context)
	DeleteOne(c *gin.Context)
	DeleteMore(c *gin.Context)
}
