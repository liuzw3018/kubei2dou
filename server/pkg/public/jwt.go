package public

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/liuzw3018/saber/lib"
	"time"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/26 15:06
 * @File: jwt.go
 * @Description: //TODO $
 * @Version:
 */

type Users struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomClaims struct {
	Users
	jwt.StandardClaims
}

var MySecret = []byte("otto_lzn")

// GenToken 创建 Token
func GenToken(user Users) (string, error) {
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() +
				int64(lib.GetIntConf("base.jwt.expires_at")), //5分钟后过期
			Issuer:    "otto",                 //签发人
			NotBefore: time.Now().Unix() - 60, // 60秒前生效
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

// ParseToken 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		//fmt.Println(" token parse err:", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新 Token
func RefreshToken(tokenStr string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Unix() + 60*10
		return GenToken(claims.Users)
	}
	return "", errors.New("Cloudn't handle this token")
}
