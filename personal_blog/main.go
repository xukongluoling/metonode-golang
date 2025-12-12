package main

import (
	"log"

	"metonode-golang/personal_blog/controllers"
	"metonode-golang/personal_blog/database"
	"metonode-golang/personal_blog/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库连接
	database.InitMySqlDB()
	log.Println("数据库连接成功")

	// 创建控制器实例
	userController := controllers.NewUserController()

	// 创建路由
	r := gin.Default()

	// 公开路由（无需认证）
	blogGop := r.Group("/blog/api")

	blogGop.POST("/register", userController.RegisterGin)
	blogGop.POST("/login", userController.LoginGin)

	// 创建控制器实例
	postController := controllers.NewPostController()
	commentController := controllers.NewCommentController()

	// 受保护路由（需要认证）
	protected := blogGop.Group("/")
	protected.Use(middleware.AuthMiddlewareGin())

	// 文章相关路由
	protected.GET("/posts", postController.GetAllPostsGin)
	protected.POST("/posts/create", postController.CreatePostGin)
	protected.PUT("/posts/update/:id", postController.UpdatePostGin)
	protected.DELETE("/posts/delete/:id", postController.DeletePostGin)
	protected.GET("/posts/:id", postController.GetPostByIDGin)

	// 评论相关路由（需要认证）
	protected.GET("/comments/:postId", commentController.GetCommentsByPostIDGin)
	protected.POST("/comments/create", commentController.CreateCommentGin)
	protected.PUT("/comments/update/:id", commentController.UpdateCommentGin)
	protected.DELETE("/comments/delete/:id", commentController.DeleteCommentGin)

	// 启动服务器
	log.Println("服务器正在运行，监听端口 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
