package middleware

import (
	"net/http"
	"strings"

	"metonode-golang/personal_blog/global_exceptions"
	"metonode-golang/personal_blog/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware JWT认证中间件，类似Java的切面 (兼容标准库)
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从请求头中获取Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			global_exceptions.HandlerError(w, global_exceptions.ErrUnauthorized)
			return
		}

		// 检查Authorization头格式是否为Bearer token
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			global_exceptions.HandlerError(w, global_exceptions.ErrUnauthorized)
			return
		}

		// 解析和验证token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			global_exceptions.HandlerError(w, global_exceptions.ErrUnauthorized)
			return
		}

		// 将用户信息存储到请求上下文（可选，方便后续处理函数使用）
		r = r.WithContext(utils.SetUserInfo(r.Context(), claims.UserID, claims.Username))

		// 继续处理请求
		next(w, r)
	}
}

// AuthMiddlewareGin JWT认证中间件 (Gin版本)
func AuthMiddlewareGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
			c.Abort()
			return
		}

		// 检查Authorization头格式是否为Bearer token
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
			c.Abort()
			return
		}

		// 解析和验证token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
			c.Abort()
			return
		}

		// 将用户信息存储到Gin上下文（方便后续处理函数使用）
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		// 继续处理请求
		c.Next()
	}
}
