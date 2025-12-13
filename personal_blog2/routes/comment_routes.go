package routes

import (
	"metonode-golang/personal_blog2/controllers"

	"github.com/gin-gonic/gin"
)

func setupCommentRoutes(router *gin.RouterGroup, commentController *controllers.CommentController) {
	// 公开路由
	router.GET("/comments/:postId", commentController.GetCommentsByPostIDGin)

	// 需要认证的路由
	protected := setupProtectedRoutes(router)
	protected.POST("/comments/create", commentController.CreateCommentGin)
	protected.PUT("/comments/update/:id", commentController.UpdateCommentGin)
	protected.DELETE("/comments/delete/:id", commentController.DeleteCommentGin)

}
