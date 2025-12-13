package services

import (
	"errors"
	"metonode-golang/personal_blog2/models"
	"metonode-golang/personal_blog2/repositories"
	"metonode-golang/personal_blog2/utils"
)

type UserService interface {
	RegisterUser(user *models.User) error
	LoginUser(username, password string) (string, error)
	GetUserByID(id uint) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService() UserService {
	return &userService{userRepo: repositories.NewUserRepository()}
}

func (s *userService) RegisterUser(user *models.User) error {
	// 检查用户名是否存在
	userName, _ := s.userRepo.GetUserByUsername(user.Username)
	if userName != nil {
		return errors.New("user already exists")
	}

	// 验证用户输入
	if err := user.ValidateUser(); err != nil {
		return err
	}

	// 密码加密
	if err := user.HashPassword(); err != nil {
		return err
	}

	// 保存到数据库
	return s.userRepo.CreateUser(user)
}

func (s *userService) LoginUser(username, password string) (string, error) {
	// 用户名查找用户
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}
	// 验证密码
	if err := user.CheckPassword(password); err != nil {
		return "", errors.New("用户名或密码错误")
	}
	// 生成token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", err
	}
	return token, nil

}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}
