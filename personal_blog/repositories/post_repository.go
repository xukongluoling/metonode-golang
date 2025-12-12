package repositories

import (
	"metonode-golang/personal_blog/database"
	"metonode-golang/personal_blog/models"

	"gorm.io/gorm"
)

// PostRepository 文章数据访问层
type PostRepository struct {
	DB *gorm.DB
}

// NewPostRepository 创建文章数据访问层实例
func NewPostRepository() *PostRepository {
	return &PostRepository{
		DB: database.MySqlDB,
	}
}

// CreatePost 创建新文章
func (r *PostRepository) CreatePost(post *models.Post) error {
	return r.DB.Create(post).Error
}

// GetAllPosts 获取所有文章
func (r *PostRepository) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	err := r.DB.Preload("User").Find(&posts).Error
	return posts, err
}

// GetPostByID 根据ID获取文章
func (r *PostRepository) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	err := r.DB.Preload("User").Preload("Comments").Preload("Comments.User").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// UpdatePost 更新文章
func (r *PostRepository) UpdatePost(post *models.Post) error {
	return r.DB.Save(post).Error
}

// DeletePost 删除文章
func (r *PostRepository) DeletePost(id uint) error {
	return r.DB.Delete(&models.Post{}, id).Error
}
