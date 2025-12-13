package routes

import (
	"metonode-golang/personal_blog2/controllers"

	"github.com/gin-gonic/gin"
)

// setupUserRoutes 配置用户相关路由
func setupUserRoutes(routerGroup *gin.RouterGroup, userController *controllers.UserController) {
	// 公开路由（无需认证）
	routerGroup.POST("/register", userController.Register)
	routerGroup.POST("/login", userController.Login)
}
