package global_exceptions

import "net/http"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

// 常用错误码
var (
	ErrBadRequest     = &AppError{Code: http.StatusBadRequest, Message: "Bad Request"}
	ErrNotFound       = &AppError{Code: http.StatusNotFound, Message: "Resource Not Found"}
	ErrInternalServer = &AppError{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
	ErrUnauthorized   = &AppError{Code: http.StatusUnauthorized, Message: "Unauthorized"}
	ErrForbidden      = &AppError{Code: http.StatusForbidden, Message: "Forbidden"}
)

// NewAppError 从错误创建 AppError
func NewAppError(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

// WrapError 从标准错误创建 AppError
func WrapError(err error, code int) *AppError {
	if err == nil {
		return nil
	}
	return &AppError{Code: code, Message: err.Error()}
}
