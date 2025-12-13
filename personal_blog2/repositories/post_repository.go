package repositories

import (
	"metonode-golang/personal_blog2/database"
	"metonode-golang/personal_blog2/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(post *models.Post) error
	GetAllPosts() ([]models.Post, error)
	GetPostByID(id uint) (*models.Post, error)
	UpdatePost(post *models.Post) error
	DeletePost(id uint) error
}

type postRepository struct {
	DB *gorm.DB
}

func NewPostRepository() PostRepository {
	return &postRepository{DB: database.MysqlDB}
}

// CreatePost 新增文章
func (pr *postRepository) CreatePost(post *models.Post) error {
	return pr.DB.Create(post).Error
}

// GetAllPosts 获取所有文章
func (pr *postRepository) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	err := pr.DB.Preload("User").Find(&posts).Error
	return posts, err
}

// GetPostByID id获取文章
func (pr *postRepository) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	err := pr.DB.Preload("User").First(&post, id).Error
	return &post, err
}

// UpdatePost 更新文章
func (pr *postRepository) UpdatePost(post *models.Post) error {
	return pr.DB.Save(post).Error
}

// DeletePost 删除文章
func (pr *postRepository) DeletePost(id uint) error {
	return pr.DB.Delete(&models.Post{}, id).Error
}
