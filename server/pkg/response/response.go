package response

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/kubei2dou/server/pkg/public"
	"github.com/liuzw3018/saber/lib"
	"strings"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/25 17:56
 * @File: response.go
 * @Description: //TODO $
 * @Version:
 */

type Response struct {
	Code    public.ResponseCode `json:"code"`
	Data    interface{}         `json:"data"`
	Msg     string              `json:"message"`
	TraceId string              `json:"trace_id"`
	Stack   string              `json:"stack"`
}

func Error(c *gin.Context, code public.ResponseCode, err error) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*lib.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	stack := ""
	if c.Query("is_debug") == "1" || lib.GetConfEnv() == "dev" {
		stack = strings.Replace(fmt.Sprintf("%+v", err), err.Error()+"\n", "", -1)
	}

	resp := &Response{Code: code, Msg: err.Error(), Data: "", TraceId: traceId, Stack: stack}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200, err)
}

func Success(c *gin.Context, code public.ResponseCode, message string, data interface{}) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*lib.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	resp := &Response{Code: code, Msg: message, Data: data, TraceId: traceId}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
