package services

import (
	"metonode-golang/personal_blog2/constants"
	"metonode-golang/personal_blog2/controllers/dto"
	"metonode-golang/personal_blog2/models"
	"metonode-golang/personal_blog2/repositories"
)

type PostService interface {
	CreatePost(reqDto *dto.CreatePostRequest) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
	GetPostByID(id uint) (*models.Post, error)
	UpdatePost(reqDto *dto.UpdatePostRequest) (*models.Post, error)
	DeletePost(id uint, userID uint) error
}

type postService struct {
	postRepo repositories.PostRepository
}

func NewPostService() PostService {
	return &postService{postRepo: repositories.NewPostRepository()}
}

// CreatePost 创建文章
func (s *postService) CreatePost(reqDto *dto.CreatePostRequest) (*models.Post, error) {
	post := reqDto.ReqToModel()

	// 文章输入内容验证
	if err := post.ValidatePost(); err != nil {
		return nil, err
	}

	// 保存到数据库
	if err := s.postRepo.CreatePost(post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *postService) GetAllPosts() ([]models.Post, error) {
	return s.postRepo.GetAllPosts()
}

func (s *postService) GetPostByID(id uint) (*models.Post, error) {
	return s.postRepo.GetPostByID(id)
}

func (s *postService) UpdatePost(reqDto *dto.UpdatePostRequest) (*models.Post, error) {
	// 获取文章
	post, err := s.postRepo.GetPostByID(reqDto.Id)
	if err != nil {
		return nil, err
	}

	// 判断是否是作者
	if post.UserID != reqDto.UserID {
		return nil, constants.ErrNotAuthor
	}
	post.Title = reqDto.Title
	post.Content = reqDto.Content
	if err := s.postRepo.UpdatePost(post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *postService) DeletePost(id uint, userID uint) error {
	// 获取文章
	post, err := s.postRepo.GetPostByID(id)
	if err != nil {
		return err
	}

	// 判断是否是作者
	if post.UserID != userID {
		return constants.ErrNotAuthor
	}

	// 删除文章
	return s.postRepo.DeletePost(id)
}
