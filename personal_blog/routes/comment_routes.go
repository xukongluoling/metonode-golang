package routes

import (
	"metonode-golang/personal_blog/controllers"

	"github.com/gin-gonic/gin"
)

// setupCommentRoutes 配置评论相关路由
func setupCommentRoutes(routerGroup *gin.RouterGroup, commentController *controllers.CommentController) {
	// 公开路由
	routerGroup.GET("/comments/:postId", commentController.GetCommentsByPostIDGin)
	
	// 需要认证的路由
	protected := setupProtectedRoutes(routerGroup)
	protected.POST("/comments/create", commentController.CreateCommentGin)
	protected.PUT("/comments/update/:id", commentController.UpdateCommentGin)
	protected.DELETE("/comments/delete/:id", commentController.DeleteCommentGin)
}
