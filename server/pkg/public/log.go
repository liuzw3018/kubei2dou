package public

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/saber/lib"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/26 13:07
 * @File: log.go
 * @Description: //TODO $
 * @Version:
 */

// ContextWarning 错误日志
func ContextWarning(c context.Context, srtag string, m map[string]interface{}) {
	v := c.Value("trace")
	traceContext, ok := v.(*lib.TraceContext)
	if !ok {
		traceContext = lib.NewTrace()
	}
	lib.Log.TagWarn(traceContext, srtag, m)
}

// ContextError 错误日志
func ContextError(c context.Context, srtag string, m map[string]interface{}) {
	v := c.Value("trace")
	traceContext, ok := v.(*lib.TraceContext)
	if !ok {
		traceContext = lib.NewTrace()
	}
	lib.Log.TagError(traceContext, srtag, m)
}

// ContextNotice 普通日志
func ContextNotice(c context.Context, srtag string, m map[string]interface{}) {
	v := c.Value("trace")
	traceContext, ok := v.(*lib.TraceContext)
	if !ok {
		traceContext = lib.NewTrace()
	}
	lib.Log.TagInfo(traceContext, srtag, m)
}

// ComLogWarning 错误日志
func ComLogWarning(c *gin.Context, srtag string, m map[string]interface{}) {
	traceContext := GetGinTraceContext(c)
	lib.Log.TagError(traceContext, srtag, m)
}

// ComLogNotice 普通日志
func ComLogNotice(c *gin.Context, srtag string, m map[string]interface{}) {
	traceContext := GetGinTraceContext(c)
	lib.Log.TagInfo(traceContext, srtag, m)
}

// GetGinTraceContext 从gin的Context中获取数据
func GetGinTraceContext(c *gin.Context) *lib.TraceContext {
	// 防御
	if c == nil {
		return lib.NewTrace()
	}
	traceContext, exists := c.Get("trace")
	if exists {
		if tc, ok := traceContext.(*lib.TraceContext); ok {
			return tc
		}
	}
	return lib.NewTrace()
}

// GetTraceContext 从Context中获取数据
func GetTraceContext(c context.Context) *lib.TraceContext {
	if c == nil {
		return lib.NewTrace()
	}
	traceContext := c.Value("trace")
	if tc, ok := traceContext.(*lib.TraceContext); ok {
		return tc
	}
	return lib.NewTrace()
}
