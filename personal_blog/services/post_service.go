package services

import (
	"metonode-golang/personal_blog/models"
	"metonode-golang/personal_blog/repositories"
)

// PostService 文章服务层
type PostService struct {
	postRepo *repositories.PostRepository
}

// NewPostService 创建文章服务层实例
func NewPostService() *PostService {
	return &PostService{
		postRepo: repositories.NewPostRepository(),
	}
}

// CreatePost 创建新文章
func (s *PostService) CreatePost(title, content string, userID uint) (*models.Post, error) {
	post := &models.Post{
		Title:   title,
		Content: content,
		UserID:  userID,
	}

	// 验证文章输入
	if err := post.ValidatePost(); err != nil {
		return nil, err
	}

	// 保存文章到数据库
	if err := s.postRepo.CreatePost(post); err != nil {
		return nil, err
	}

	return post, nil
}

// GetAllPosts 获取所有文章
func (s *PostService) GetAllPosts() ([]models.Post, error) {
	return s.postRepo.GetAllPosts()
}

// GetPostByID 根据ID获取文章
func (s *PostService) GetPostByID(id uint) (*models.Post, error) {
	return s.postRepo.GetPostByID(id)
}

// UpdatePost 更新文章
func (s *PostService) UpdatePost(id uint, title, content string, userID uint) (*models.Post, error) {
	// 获取文章
	post, err := s.postRepo.GetPostByID(id)
	if err != nil {
		return nil, err
	}

	// 检查是否是文章作者
	if post.UserID != userID {
		return nil, models.ErrNotAuthor
	}

	// 更新文章内容
	post.Title = title
	post.Content = content

	// 保存更新
	if err = s.postRepo.UpdatePost(post); err != nil {
		return nil, err
	}

	return post, nil
}

// DeletePost 删除文章
func (s *PostService) DeletePost(id uint, userID uint) error {
	// 获取文章
	post, err := s.postRepo.GetPostByID(id)
	if err != nil {
		return err
	}

	// 检查是否是文章作者
	if post.UserID != userID {
		return models.ErrNotAuthor
	}

	// 删除文章
	return s.postRepo.DeletePost(id)
}
