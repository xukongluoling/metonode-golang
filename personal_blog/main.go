package main

import (
	"log"

	"metonode-golang/personal_blog/config"
	"metonode-golang/personal_blog/database"
	"metonode-golang/personal_blog/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}
	log.Println("配置文件加载成功")

	// 初始化数据库连接
	database.InitMySqlDB()
	log.Println("数据库连接成功")

	// 创建路由
	r := gin.Default()

	// 配置所有路由
	routes.SetupRouter(r)

	// 启动服务器
	log.Println("服务器正在运行，监听端口 8080...")
	if err := r.Run(config.AppConfig.Server.Port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
