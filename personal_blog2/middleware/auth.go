package middleware

import (
	"metonode-golang/personal_blog2/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求入参获取Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is empty"})
			c.Abort()
			return
		}

		// 检查Authorization值是否为Bearer token格式
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is invalid"})
			c.Abort()
			return
		}

		// 解析验证token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is invalid"})
			c.Abort()
			return
		}
		// 用户信息存到Gin上下文
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		// 继续处理请求
		c.Next()
	}
}
