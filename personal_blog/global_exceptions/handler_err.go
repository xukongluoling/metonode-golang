package global_exceptions

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// ErrorHandler 适配标准http.HandlerFunc的ErrorHandler
func ErrorHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v\n%s", err, string(debug.Stack()))
			}
		}()
		// 执行下一个 handler
		next(w, req)
	}
}

// GinErrorHandler 适配Gin框架的ErrorHandler
func GinErrorHandler(next http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v\n%s", err, string(debug.Stack()))
			}
		}()
		// 执行下一个 handler
		next(c.Writer, c.Request)
	}
}
func handlerError(w http.ResponseWriter, err *AppError) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(err.Code)
	_, err2 := w.Write([]byte(fmt.Sprintf("{\"code\": %d, \"message\": %s}", err.Code, err.Message)))
	if err2 != nil {
		return
	}
}

func HandlerError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	// 检查是否是AppError
	var appErr *AppError
	if errors.As(err, &appErr) {
		handlerError(w, appErr)
		return
	}
	log.Printf("Unknown error: %v", err)
	handlerError(w, ErrInternalServer)
}
