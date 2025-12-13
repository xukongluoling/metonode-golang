package controllers

import (
	"metonode-golang/personal_blog2/controllers/dto"
	"metonode-golang/personal_blog2/services"
	"metonode-golang/personal_blog2/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	userService services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	var request dto.RegisterRequest

	if !utils.BindAndValidate(ctx, &request) {
		utils.Logger.Info("用户注册参数验证失败", zap.String("username", request.Username))
		return
	}

	utils.Logger.Info("用户注册请求", zap.String("username", request.Username), zap.String("email", request.Email))

	user := request.RegisterToUser()
	if err := c.userService.RegisterUser(user); err != nil {
		utils.Logger.Error("用户注册失败", zap.String("username", request.Username), zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.Logger.Info("用户注册成功", zap.String("username", request.Username))
	ctx.JSON(http.StatusCreated, gin.H{"message": "注册成功"})
}

func (c *UserController) Login(ctx *gin.Context) {
	var request dto.LoginRequest

	if !utils.BindAndValidate(ctx, &request) {
		utils.Logger.Info("用户登录参数验证失败", zap.String("username", request.Username))
		return
	}

	utils.Logger.Info("用户登录请求", zap.String("username", request.Username))

	// 登录逻辑
	token, err := c.userService.LoginUser(request.Username, request.Password)
	if err != nil {
		utils.Logger.Error("用户登录失败", zap.String("username", request.Username), zap.Error(err))
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	utils.Logger.Info("用户登录成功", zap.String("username", request.Username))
	ctx.JSON(http.StatusOK, gin.H{"message": "登录成功", "token": token})
}
