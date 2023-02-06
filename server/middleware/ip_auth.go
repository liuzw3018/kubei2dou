package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/kubei2dou/server/pkg/public"
	"github.com/liuzw3018/kubei2dou/server/pkg/response"
	"github.com/liuzw3018/saber/lib"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/26 13:14
 * @File: ip_auth.go
 * @Description: //TODO $
 * @Version:
 */

// IPAuthMiddleware 白名单验证
func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isMatched := false
		for _, host := range lib.GetStringSliceConf("base.http.allow_ip") {
			if c.ClientIP() == host {
				isMatched = true
			}
		}
		if !isMatched {
			response.Error(c, public.InternalErrorCode, errors.New(fmt.Sprintf("%v, not in iplist", c.ClientIP())))
			c.Abort()
			return
		}
		c.Next()
	}
}
