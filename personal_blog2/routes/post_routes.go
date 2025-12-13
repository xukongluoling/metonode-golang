package routes

import (
	"metonode-golang/personal_blog2/controllers"

	"github.com/gin-gonic/gin"
)

func setPostRoutes(rg *gin.RouterGroup, postController *controllers.PostController) {
	// 公开路由
	rg.GET("/posts/:id", postController.GetPostByIDGin)

	// 需要认证的路由
	protected := setupProtectedRoutes(rg)
	protected.GET("/posts", postController.GetAllPost)
	protected.POST("/posts/create", postController.CreatePost)
	protected.PUT("/posts/update/:id", postController.UpdatePost)
	protected.DELETE("/posts/delete/:id", postController.DeletePost)
}
