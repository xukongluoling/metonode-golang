package services

import (
	"errors"

	"metonode-golang/personal_blog/models"
	"metonode-golang/personal_blog/repositories"
	"metonode-golang/personal_blog/utils"
)

// UserService 用户服务层
type UserService struct {
	userRepo *repositories.UserRepository
}

// NewUserService 创建用户服务层实例
func NewUserService() *UserService {
	return &UserService{
		userRepo: repositories.NewUserRepository(),
	}
}

// RegisterUser 用户注册
func (s *UserService) RegisterUser(username, password, email string) error {
	// 检查用户名是否已存在
	existingUser, _ := s.userRepo.GetUserByUsername(username)
	if existingUser != nil {
		return errors.New("用户名已存在")
	}

	// 创建新用户
	user := &models.User{
		Username: username,
		Password: password,
		Email:    email,
	}

	// 验证用户输入
	if err := user.ValidateUser(); err != nil {
		return err
	}

	// 密码加密
	if err := user.HashPassword(); err != nil {
		return err
	}

	// 保存用户到数据库
	return s.userRepo.CreateUser(user)
}

// LoginUser 用户登录，返回JWT token
func (s *UserService) LoginUser(username, password string) (string, error) {
	// 根据用户名查找用户
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	// 验证密码
	if err := user.CheckPassword(password); err != nil {
		return "", errors.New("用户名或密码错误")
	}

	// 生成JWT token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserByID 根据ID获取用户信息
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}
