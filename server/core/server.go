package core

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/kubei2dou/server/initialize"
	_ "github.com/liuzw3018/kubei2dou/server/pkg/public"
	"github.com/liuzw3018/saber/lib"
	"log"
	"net/http"
	"time"
)

/**
 * @Author: liu zw
 * @Date: 2022/9/22 11:48
 * @File: server.go
 * @Description: //TODO $
 * @Version:
 */

var (
	srv *http.Server
)

func RunServer() {
	gin.SetMode(lib.ConfBase.DebugMode)
	r := initialize.InitRouters()
	srv = initHttpServer(r)
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", lib.GetStringConf("base.http.addr"))
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", lib.GetStringConf("base.http.addr"), err)
		}
	}()
}

func StopServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
