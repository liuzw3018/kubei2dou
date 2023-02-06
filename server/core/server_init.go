package core

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/saber/lib"
	"net/http"
	"time"
)

/**
 * @Author: liu zw
 * @Date: 2022/9/22 11:54
 * @File: server_init.go
 * @Description: //TODO $
 * @Version:
 */

func initHttpServer(handlerFunc *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           lib.GetStringConf("base.http.addr"),
		Handler:        handlerFunc,
		ReadTimeout:    time.Duration(lib.GetIntConf("base.http.read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(lib.GetIntConf("base.http.write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << uint(lib.GetIntConf("base.http.max_header_bytes")),
	}
}
