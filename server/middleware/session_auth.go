package middleware

import (
	"errors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/kubei2dou/server/pkg/public"
	"github.com/liuzw3018/kubei2dou/server/pkg/response"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/26 13:59
 * @File: session_auth.go
 * @Description: //TODO $
 * @Version:
 */

// SessionAuthMiddleware Session验证
func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if name, ok := session.Get("user").(string); !ok || name == "" {
			response.Error(c, public.InternalErrorCode, errors.New("user not login"))
			c.Abort()
			return
		}
		c.Next()
	}
}
