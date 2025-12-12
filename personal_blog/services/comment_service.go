package services

import (
	"metonode-golang/personal_blog/models"
	"metonode-golang/personal_blog/repositories"
)

// CommentService 评论服务层
type CommentService struct {
	commentRepo *repositories.CommentRepository
	postRepo    *repositories.PostRepository
}

// NewCommentService 创建评论服务层实例
func NewCommentService() *CommentService {
	return &CommentService{
		commentRepo: repositories.NewCommentRepository(),
		postRepo:    repositories.NewPostRepository(),
	}
}

// CreateComment 创建新评论
func (s *CommentService) CreateComment(content string, postID, userID uint) (*models.Comment, error) {
	// 检查文章是否存在
	_, err := s.postRepo.GetPostByID(postID)
	if err != nil {
		return nil, err
	}

	comment := &models.Comment{
		Content: content,
		PostID:  postID,
		UserID:  userID,
	}

	// 验证评论输入
	if err := comment.ValidateComment(); err != nil {
		return nil, err
	}

	// 保存评论到数据库
	if err := s.commentRepo.CreateComment(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

// GetCommentsByPostID 根据文章ID获取所有评论
func (s *CommentService) GetCommentsByPostID(postID uint) ([]models.Comment, error) {
	// 检查文章是否存在
	_, err := s.postRepo.GetPostByID(postID)
	if err != nil {
		return nil, err
	}

	return s.commentRepo.GetCommentsByPostID(postID)
}

// UpdateComment 更新评论
func (s *CommentService) UpdateComment(id uint, content string, userID uint) (*models.Comment, error) {
	// 获取评论
	comment, err := s.commentRepo.GetCommentByID(id)
	if err != nil {
		return nil, err
	}

	// 检查是否是评论作者
	if comment.UserID != userID {
		return nil, models.ErrNotAuthor
	}

	// 更新评论内容
	comment.Content = content

	// 保存更新
	if err := s.commentRepo.UpdateComment(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

// DeleteComment 删除评论
func (s *CommentService) DeleteComment(id uint, userID uint) error {
	// 获取评论
	comment, err := s.commentRepo.GetCommentByID(id)
	if err != nil {
		return err
	}

	// 检查是否是评论作者
	if comment.UserID != userID {
		return models.ErrNotAuthor
	}

	// 删除评论
	return s.commentRepo.DeleteComment(id)
}
