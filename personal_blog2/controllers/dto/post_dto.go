package dto

import "metonode-golang/personal_blog2/models"

// CreatePostRequest 创建文章请求结构体
type CreatePostRequest struct {
	UserID  uint   `json:"user_id"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

// UpdatePostRequest 更新文章请求结构体
type UpdatePostRequest struct {
	Id      uint   `json:"id" validate:"required"`
	UserID  uint   `json:"userId" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (req *CreatePostRequest) ReqToModel() *models.Post {
	return &models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}
}

func (req *UpdatePostRequest) ReqToModel(userID uint) *models.Post {
	return &models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}
}
