package repositories

import (
	"metonode-golang/personal_blog2/database"
	"metonode-golang/personal_blog2/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(comment *models.Comment) error
	GetCommentsByPostID(postID uint) ([]models.Comment, error)
	GetCommentByID(id uint) (*models.Comment, error)
	UpdateComment(comment *models.Comment) error
	DeleteComment(id uint) error
}

type commentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository() CommentRepository {
	return &commentRepository{DB: database.MysqlDB}
}

// CreateComment 新增评论
func (r *commentRepository) CreateComment(comment *models.Comment) error {
	return r.DB.Create(comment).Error
}

// GetCommentsByPostID 获取文章所有评论
func (r *commentRepository) GetCommentsByPostID(postID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.DB.Where("post_id = ?", postID).Preload("User").Find(&comments).Error
	return comments, err
}

// GetCommentByID 根据id获取评论
func (r *commentRepository) GetCommentByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := r.DB.Where("id = ?", id).Preload("User").First(&comment).Error
	return &comment, err
}

// UpdateComment 更新评论
func (r *commentRepository) UpdateComment(comment *models.Comment) error {
	return r.DB.Save(comment).Error
}

// DeleteComment 删除评论
func (r *commentRepository) DeleteComment(id uint) error {
	return r.DB.Delete(&models.Comment{}, id).Error
}
