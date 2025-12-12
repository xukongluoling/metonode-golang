package routes

import (
	"metonode-golang/personal_blog/controllers"
	"metonode-golang/personal_blog/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 配置所有路由
func SetupRouter(r *gin.Engine) {
	// 创建控制器实例
	userController := controllers.NewUserController()
	postController := controllers.NewPostController()
	commentController := controllers.NewCommentController()

	// 基础API路由组
	apiGroup := r.Group("/blog/api")

	// 设置用户相关路由
	setupUserRoutes(apiGroup, userController)

	// 设置文章相关路由
	setupPostRoutes(apiGroup, postController)

	// 设置评论相关路由
	setupCommentRoutes(apiGroup, commentController)
}

// setupProtectedRoutes 创建需要认证的路由组
func setupProtectedRoutes(parent *gin.RouterGroup) *gin.RouterGroup {
	return parent.Group("/").Use(middleware.AuthMiddlewareGin()).(*gin.RouterGroup)
}
