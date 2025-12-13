package dto

import "metonode-golang/personal_blog2/models"

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

func (r *RegisterRequest) RegisterToUser() *models.User {
	return &models.User{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}
}

//func (r *LoginRequest) LoginToUser() *models.User {
//	return &models.User{
//		Username: r.Username,
//		Password: r.Password,
//	}
//}
