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
	val := validator.New()
	return val.Struct(u)
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
	val := validator.New()
	return val.Struct(p)
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
	PostID  uint
	Post    Post `gorm:"foreignKey:PostID"`
}

func (c *Comment) ValidateComment() error {
	if c.Content == "" {
		return errors.New("评论内容不能为空")
	}
	return nil
}

// HashPassword 密码哈希
func (u *User) HashPassword() error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}

// CheckPassword 密码验证
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
