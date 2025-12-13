package services

import (
	"metonode-golang/personal_blog2/constants"
	"metonode-golang/personal_blog2/models"
	"metonode-golang/personal_blog2/repositories"
)

type CommentService interface {
	CreateComment(content string, postID, userID uint) (*models.Comment, error)
	GetCommentsByPostID(postID uint) ([]models.Comment, error)
	UpdateComment(id uint, content string, userID uint) (*models.Comment, error)
	DeleteComment(id uint, userID uint) error
}

type commentService struct {
	postRepo repositories.PostRepository
	comRepo  repositories.CommentRepository
}

func NewCommentService() CommentService {
	return &commentService{
		postRepo: repositories.NewPostRepository(),
		comRepo:  repositories.NewCommentRepository()}
}

func (s *commentService) CreateComment(content string, postID, userID uint) (*models.Comment, error) {
	// 检查文章是否存在
	_, err := s.postRepo.GetPostByID(postID)
	if err != nil {
		return nil, err
	}

	comment := &models.Comment{
		PostID:  postID,
		UserID:  userID,
		Content: content,
	}

	// 检验数据
	if err := comment.ValidateComment(); err != nil {
		return nil, err
	}

	if err := s.comRepo.CreateComment(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *commentService) GetCommentsByPostID(postID uint) ([]models.Comment, error) {
	// 获取文章
	_, err := s.postRepo.GetPostByID(postID)
	if err != nil {
		return nil, err
	}
	return s.comRepo.GetCommentsByPostID(postID)
}

func (s *commentService) UpdateComment(id uint, content string, userID uint) (*models.Comment, error) {
	// 获取评论
	comment, err := s.comRepo.GetCommentByID(id)
	if err != nil {
		return nil, err
	}
	// 是否是作者
	if comment.UserID != userID {
		return nil, constants.ErrNotAuthor
	}
	comment.Content = content

	if err := s.comRepo.UpdateComment(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *commentService) DeleteComment(id uint, userID uint) error {
	// 获取评论
	comment, err := s.comRepo.GetCommentByID(id)
	if err != nil {
		return err
	}
	// 是否作者
	if comment.UserID != userID {
		return constants.ErrNotAuthor
	}
	return s.comRepo.DeleteComment(id)
}
