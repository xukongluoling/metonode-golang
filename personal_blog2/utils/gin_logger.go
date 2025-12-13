package utils

import (
	"io"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
) // GinLogger 自定义Gin日志中间件，将日志输出到zap
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		end := time.Now()
		// 计算处理时间
		duration := end.Sub(start)

		// 请求方式
		method := c.Request.Method
		// 请求路由
		path := c.Request.URL.Path
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 错误信息
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
		// 请求大小
		requestSize := c.Request.ContentLength
		// 响应大小
		responseSize := c.Writer.Size()

		// 记录日志
		if statusCode >= 500 {
			// 服务器错误
			Logger.Error("Gin请求处理失败",
				zap.Int("status_code", statusCode),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("client_ip", clientIP),
				zap.String("error", errorMessage),
				zap.Duration("duration", duration),
				zap.Int64("request_size", requestSize),
				zap.Int("response_size", responseSize),
			)
		} else if statusCode >= 400 {
			// 客户端错误
			Logger.Warn("Gin请求客户端错误",
				zap.Int("status_code", statusCode),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("client_ip", clientIP),
				zap.String("error", errorMessage),
				zap.Duration("duration", duration),
				zap.Int64("request_size", requestSize),
				zap.Int("response_size", responseSize),
			)
		} else {
			// 正常请求
			Logger.Info("Gin请求处理完成",
				zap.Int("status_code", statusCode),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("client_ip", clientIP),
				zap.Duration("duration", duration),
				zap.Int64("request_size", requestSize),
				zap.Int("response_size", responseSize),
			)
		}
	}
}

// GinRecovery 自定义Gin恢复中间件，将异常日志输出到zap
func GinRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		var errMsg string
		var stackTrace string

		// 尝试将recovered转换为error类型
		switch r := recovered.(type) {
		case error:
			errMsg = r.Error()
			stackTrace = string(debug.Stack())
		case string:
			errMsg = r
			stackTrace = string(debug.Stack())
		default:
			errMsg = "unknown panic"
			stackTrace = string(debug.Stack())
		}

		// 记录错误日志
		Logger.Error("Gin请求崩溃",
			zap.String("error", errMsg),
			zap.String("stack_trace", stackTrace),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("client_ip", c.ClientIP()),
		)

		// 返回500错误
		c.AbortWithStatusJSON(500, gin.H{
			"error": "Internal Server Error",
		})
	})
}

// DisableGinDefaultLog 禁用Gin默认日志
func DisableGinDefaultLog(r *gin.Engine) {
	// 禁用默认的Logger中间件
	r.Use(func(c *gin.Context) {
		c.Next()
	})

	// 将Gin的默认输出重定向到空
	gin.DefaultWriter = io.Discard
	gin.DefaultErrorWriter = io.Discard

	// 禁用控制台颜色
	gin.DisableConsoleColor()
}
