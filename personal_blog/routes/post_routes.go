package routes

import (
	"metonode-golang/personal_blog/controllers"

	"github.com/gin-gonic/gin"
)

// setupPostRoutes 配置文章相关路由
func setupPostRoutes(routerGroup *gin.RouterGroup, postController *controllers.PostController) {
	// 公开路由
	routerGroup.GET("/posts/:id", postController.GetPostByIDGin)

	// 需要认证的路由
	protected := setupProtectedRoutes(routerGroup)
	protected.GET("/posts", postController.GetAllPostsGin)
	protected.POST("/posts/create", postController.CreatePostGin)
	protected.PUT("/posts/update/:id", postController.UpdatePostGin)
	protected.DELETE("/posts/delete/:id", postController.DeletePostGin)
}
