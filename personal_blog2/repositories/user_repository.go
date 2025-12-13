package repositories

import (
	"metonode-golang/personal_blog2/database"
	"metonode-golang/personal_blog2/models"

	"gorm.io/gorm"
)

// UserRepository 用户repository接口
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{DB: database.MysqlDB}
}

// CreateUser 新增用户
func (r *userRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

// GetUserByUsername 用户名查询用户
func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID 用户id查询用户
func (r *userRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
