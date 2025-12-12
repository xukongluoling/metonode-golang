package repositories

import (
	"metonode-golang/personal_blog/database"
	"metonode-golang/personal_blog/models"

	"gorm.io/gorm"
)

// CommentRepository 评论数据访问层
type CommentRepository struct {
	DB *gorm.DB
}

// NewCommentRepository 创建评论数据访问层实例
func NewCommentRepository() *CommentRepository {
	return &CommentRepository{
		DB: database.MySqlDB,
	}
}

// CreateComment 创建新评论
func (r *CommentRepository) CreateComment(comment *models.Comment) error {
	return r.DB.Create(comment).Error
}

// GetCommentsByPostID 根据文章ID获取所有评论
func (r *CommentRepository) GetCommentsByPostID(postID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.DB.Where("post_id = ?", postID).Preload("User").Find(&comments).Error
	return comments, err
}

// GetCommentByID 根据ID获取评论
func (r *CommentRepository) GetCommentByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := r.DB.Preload("User").Preload("Post").First(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// UpdateComment 更新评论
func (r *CommentRepository) UpdateComment(comment *models.Comment) error {
	return r.DB.Save(comment).Error
}

// DeleteComment 删除评论
func (r *CommentRepository) DeleteComment(id uint) error {
	return r.DB.Delete(&models.Comment{}, id).Error
}
