package repositories

import (
	"metonode-golang/personal_blog/database"
	"metonode-golang/personal_blog/models"

	"gorm.io/gorm"
)

// UserRepository 用户数据访问层
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository 创建用户数据访问层实例
func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: database.MySqlDB,
	}
}

// CreateUser 创建新用户
func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

// GetUserByUsername 根据用户名获取用户
func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID 根据ID获取用户
func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
