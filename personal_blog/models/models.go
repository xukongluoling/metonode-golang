package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `grom:"not null" json:"username"`
	Password string `grom:"not null" json:"password"`
	Email    string `grom:"not null" json:"email"`
}

func (u *User) ValidateUser() error {
	validate := validator.New()
	return validate.Struct(u)
}

type Post struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	UserID   uint
	User     User      `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}

func (p *Post) ValidatePost() error {
	validate := validator.New()
	return validate.Struct(p)
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
	PostID  uint
	Post    Post `gorm:"foreignKey:PostID"`
}

// ValidateComment 验证评论输入
func (c *Comment) ValidateComment() error {
	if c.Content == "" {
		return errors.New("评论内容不能为空")
	}
	return nil
}

// HashPassword 为用户密码生成哈希值
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword 验证用户密码是否正确
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// 定义文章相关错误
var (
	ErrNotAuthor = errors.New("不是文章作者")
)
