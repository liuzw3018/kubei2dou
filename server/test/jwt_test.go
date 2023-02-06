package test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/9 10:29
 * @File: jwt_test.go
 * @Description: //TODO $
 * @Version:
 */

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func TestJWT(t *testing.T) {
	// MapClaims map
	mySigningKey := []byte("hello")
	// StandardClaims struct
	c := MyClaims{
		Username: "lzn",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    "lvg",
			NotBefore: time.Now().Unix() - 60,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedString, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
	fmt.Println(signedString)

	token1, err := jwt.ParseWithClaims(signedString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token1)
}
