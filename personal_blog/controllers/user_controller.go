package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"metonode-golang/personal_blog/global_exceptions"
	"metonode-golang/personal_blog/services"

	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct {
	userService *services.UserService
}

// NewUserController 创建用户控制器实例
func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

// RegisterRequest 注册请求结构体
type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
}

// LoginRequest 登录请求结构体
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse 登录响应结构
type LoginResponse struct {
	Token string `json:"token"`
}

// Register 用户注册 (兼容标准库)
func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		global_exceptions.HandlerError(w, global_exceptions.ErrBadRequest)
		return
	}

	if err := c.userService.RegisterUser(req.Username, req.Password, req.Email); err != nil {
		global_exceptions.HandlerError(w, global_exceptions.NewAppError(http.StatusBadRequest, err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "注册成功"})
}

// RegisterGin 用户注册 (Gin版本)
func (c *UserController) RegisterGin(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	if err := c.userService.RegisterUser(req.Username, req.Password, req.Email); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "注册成功"})
}

// Login 用户登录 (兼容标准库)
func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("进入Login方法")
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("解码请求体失败:", err)
		global_exceptions.HandlerError(w, global_exceptions.ErrBadRequest)
		return
	}
	log.Println("解码请求体成功，用户名:", req.Username)

	token, err := c.userService.LoginUser(req.Username, req.Password)
	if err != nil {
		log.Println("登录失败:", err)
		global_exceptions.HandlerError(w, global_exceptions.NewAppError(http.StatusUnauthorized, err.Error()))
		return
	}

	// 调试日志：打印生成的token
	log.Println("生成的token:", token)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// 调试日志：打印响应内容
	resp := LoginResponse{Token: token}
	log.Println("响应内容:", resp)
	// 直接使用json.Marshal生成响应
	jsonData, err := json.Marshal(resp)
	if err != nil {
		log.Println("编码响应错误:", err)
		global_exceptions.HandlerError(w, global_exceptions.NewAppError(http.StatusInternalServerError, "编码响应错误"))
		return
	}
	log.Println("响应JSON数据:", string(jsonData))
	// 使用w.Write写入响应
	if _, err := w.Write(jsonData); err != nil {
		log.Println("写入响应错误:", err)
		global_exceptions.HandlerError(w, global_exceptions.NewAppError(http.StatusInternalServerError, "写入响应错误"))
		return
	}
	log.Println("响应发送成功")
}

// LoginGin 用户登录 (Gin版本)
func (c *UserController) LoginGin(ctx *gin.Context) {
	log.Println("进入LoginGin方法")
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println("绑定请求体失败:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}
	log.Println("绑定请求体成功，用户名:", req.Username)

	token, err := c.userService.LoginUser(req.Username, req.Password)
	if err != nil {
		log.Println("登录失败:", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	log.Println("生成的token:", token)
	ctx.JSON(http.StatusOK, LoginResponse{Token: token})
}
