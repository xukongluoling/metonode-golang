package routes

import (
	"metonode-golang/personal_blog2/controllers"
	"metonode-golang/personal_blog2/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouters 配置所有路由
func SetupRouters(r *gin.Engine) {
	userController := controllers.NewUserController()
	postController := controllers.NewPostController()
	commentController := controllers.NewCommentController()

	apiGroup := r.Group("/blog/api")

	setupUserRoutes(apiGroup, userController)
	setPostRoutes(apiGroup, postController)
	setupCommentRoutes(apiGroup, commentController)
}

func setupProtectedRoutes(parent *gin.RouterGroup) *gin.RouterGroup {
	return parent.Group("/").Use(middleware.AuthMiddleware()).(*gin.RouterGroup)
}
