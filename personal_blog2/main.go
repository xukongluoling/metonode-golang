package main

import (
	"metonode-golang/personal_blog2/config"
	"metonode-golang/personal_blog2/database"
	"metonode-golang/personal_blog2/routes"
	"metonode-golang/personal_blog2/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 初始化日志
	utils.InitLogger()
	defer utils.SyncLogger()

	// 加载配置文件
	err := config.LoadConfig()
	if err != nil {
		utils.Logger.Fatal("Error loading config", zap.Error(err))
	}

	// 初始化数据库连接
	database.InitMysqlDB()

	// 创建路由（不使用默认中间件）
	r := gin.New()

	// 应用自定义中间件
	r.Use(utils.GinLogger(), utils.GinRecovery())

	// 禁用gin默认日志输出
	utils.DisableGinDefaultLog(r)

	routes.SetupRouters(r)

	// 启动服务器
	port := ":" + config.AppConfig.Server.Port
	utils.Logger.Info("服务器正在运行，监听端口", zap.String("port", port))
	if err := r.Run(port); err != nil {
		utils.Logger.Fatal("服务器启动失败", zap.Error(err))
	}

}
