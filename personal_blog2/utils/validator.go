package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 全局验证器实例
var validate = validator.New()

// BindAndValidate 绑定并验证请求参数
// 该函数封装了Gin的JSON绑定和validator的参数验证功能
// 类似于Spring Boot的@Valid注解效果
func BindAndValidate(ctx *gin.Context, obj interface{}) bool {
	// 绑定JSON参数
	if err := ctx.ShouldBindJSON(obj); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数格式错误: " + err.Error(),
		})
		return false
	}

	// 执行参数验证
	if err := validate.Struct(obj); err != nil {
		// 处理验证错误，提取友好的错误信息
		errMsg := handleValidationErrors(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errMsg,
		})
		return false
	}

	return true
}

// handleValidationErrors 处理验证错误，提取友好的错误信息
func handleValidationErrors(err error) string {
	var errMsgs []string

	// 类型断言为validator.ValidationErrors
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, e := range validationErrors {
			// 根据不同的验证规则生成友好的错误信息
			switch e.Tag() {
			case "required":
				errMsgs = append(errMsgs, fmt.Sprintf("%s字段是必填的", e.Field()))
			case "min":
				errMsgs = append(errMsgs, fmt.Sprintf("%s字段长度不能小于%d", e.Field(), e.Param()))
			case "max":
				errMsgs = append(errMsgs, fmt.Sprintf("%s字段长度不能大于%d", e.Field(), e.Param()))
			case "email":
				errMsgs = append(errMsgs, fmt.Sprintf("%s字段必须是有效的邮箱地址", e.Field()))
			case "numeric":
				errMsgs = append(errMsgs, fmt.Sprintf("%s字段必须是数字", e.Field()))
			default:
				errMsgs = append(errMsgs, fmt.Sprintf("%s字段验证失败: %s", e.Field(), e.Tag()))
			}
		}
	}

	// 将所有错误信息合并为一个字符串
	return strings.Join(errMsgs, "; ")
}
