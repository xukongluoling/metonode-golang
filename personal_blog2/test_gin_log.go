package main

import (
	"net/http"
	"time"

	"metonode-golang/personal_blog2/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 初始化日志
	utils.InitLogger()
	defer utils.SyncLogger()

	// 创建Gin路由
	r := gin.New()

	// 应用自定义日志中间件
	r.Use(utils.GinLogger(), utils.GinRecovery())
	utils.DisableGinDefaultLog(r)

	// 添加测试路由
	r.GET("/test", func(c *gin.Context) {
		// 测试正常请求
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	r.GET("/error", func(c *gin.Context) {
		// 测试错误请求
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
	})

	r.GET("/panic", func(c *gin.Context) {
		// 测试panic
		panic("test panic")
	})

	// 启动服务器
	utils.Logger.Info("测试服务器启动", zap.String("port", ":8080"))

	// 发送测试请求
	go func() {
		// 等待服务器启动
		time.Sleep(1 * time.Second)

		utils.Logger.Info("发送测试请求")

		// 发送正常请求
		resp, _ := http.Get("http://localhost:8080/test")
		resp.Body.Close()

		// 发送错误请求
		resp, _ = http.Get("http://localhost:8080/error")
		resp.Body.Close()

		// 发送panic请求
		resp, _ = http.Get("http://localhost:8080/panic")
		resp.Body.Close()

		// 关闭服务器
		time.Sleep(1 * time.Second)
		utils.Logger.Info("测试完成")
	}()

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		utils.Logger.Fatal("服务器启动失败", zap.Error(err))
	}
}
