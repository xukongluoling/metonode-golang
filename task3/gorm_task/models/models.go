package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:100;not null;unique" json:"name"`
	Password  string `gorm:"size:255;not null" json:"password"`
	PostCount int    `gorm:"default:0" json:"post_count"` // 用户文章数量统计字段
	Posts     []Post `gorm:"foreignkey:UserID" json:"posts"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Post struct {
	Id            uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title         string    `gorm:"size:200;not null" json:"title"`
	Content       string    `gorm:"type:text" json:"content"`
	UserID        uint64    `gorm:"not null" json:"user_id"`
	User          User      `gorm:"foreignkey:UserID;association_foreignkey:ID" json:"user"`
	Comments      []Comment `gorm:"foreignkey:PostID" json:"comments"`
	CommentStatus string    `gorm:"size:20;default:'无评论'" json:"comment_status"` // 评论状态字段
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Comment struct {
	Id        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	PostId    uint64 `gorm:"not null" json:"post_id"`
	UserId    uint64 `gorm:"not null" json:"user_id"`
	Username  string `gorm:"size:100;not null" json:"username"`
	Content   string `gorm:"type:text;not null" json:"content"`
	Post      Post   `gorm:"foreignkey:PostID;association_foreignkey:ID" json:"post"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// AfterCreate Post模型的AfterCreate钩子函数，在文章创建时自动更新用户的文章数量统计字段
func (p *Post) AfterCreate(tx *gorm.DB) error {
	// 更新用户的文章数量
	return tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

// AfterDelete Comment模型的AfterDelete钩子函数，在评论删除时检查文章的评论数量
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 检查文章的评论数量
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostId).Count(&count).Error; err != nil {
		return err
	}

	// 如果评论数量为0，则更新文章的评论状态为"无评论"
	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostId).Update("comment_status", "无评论").Error
	}

	return nil
}

// AfterSave Post模型的AfterCreate钩子函数，在文章创建时自动更新评论状态为"无评论"
func (p *Post) AfterSave(tx *gorm.DB) error {
	// 检查文章是否有评论
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", p.Id).Count(&count).Error; err != nil {
		return err
	}

	// 根据评论数量更新评论状态
	status := "无评论"
	if count > 0 {
		status = "有评论"
	}

	return tx.Model(p).Update("comment_status", status).Error
}
