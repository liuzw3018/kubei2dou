package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/kubei2dou/server/pkg/public"
	"net/http"
	"strings"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/26 14:01
 * @File: jwt_auth.go
 * @Description: //TODO $
 * @Version:
 */

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "无权限访问，请求未携带token",
			})
			c.Abort() //结束后续操作
			return
		}

		//按空格拆分
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}

		//解析token包含的信息
		claims, err := public.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的claims信息保存到请求的上下文c上
		c.Set("claims", claims)
		c.Next() // 后续的处理函数可以用过ctx.Get("claims")来获取当前请求的用户信息
	}
}

// CheckUserInfo 检查用户名信息
func CheckUserInfo(claims *public.CustomClaims) error {
	username := claims.Username
	password := claims.Password
	//获取数据库用户名及密码
	if username == "admin" && password == "123456" {
		return nil
	}
	return errors.New("用户名或密码错误")
}
