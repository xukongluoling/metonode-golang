package controllers

import (
	"errors"
	"metonode-golang/personal_blog2/constants"
	"metonode-golang/personal_blog2/services"
	"metonode-golang/personal_blog2/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService services.CommentService
}

func NewCommentController() *CommentController {
	return &CommentController{commentService: services.NewCommentService()}
}

// CreateCommentRequest 创建评论请求结构体
type CreateCommentRequest struct {
	Content string `json:"content" validate:"required"`
	PostID  uint   `json:"post_id" validate:"required"`
}

// UpdateCommentRequest 更新评论请求结构体
type UpdateCommentRequest struct {
	Content string `json:"content" validate:"required"`
}

// CreateCommentGin 创建评论 (Gin版本)
func (c *CommentController) CreateCommentGin(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}
	var req CreateCommentRequest
	if !utils.BindAndValidate(ctx, &req) {
		return
	}
	comment, err := c.commentService.CreateComment(req.Content, req.PostID, userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, comment)
}

// GetCommentsByPostIDGin 根据文章ID获取所有评论 (Gin版本)
func (c *CommentController) GetCommentsByPostIDGin(ctx *gin.Context) {
	postIDStr := ctx.Param("postId")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}
	comments, err := c.commentService.GetCommentsByPostID(uint(postID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "评论不存在"})
		return
	}
	ctx.JSON(http.StatusOK, comments)
}

// UpdateCommentGin 更新评论 (Gin版本)
func (c *CommentController) UpdateCommentGin(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的评论ID"})
		return
	}
	var req UpdateCommentRequest
	if !utils.BindAndValidate(ctx, &req) {
		return
	}

	comment, err := c.commentService.UpdateComment(uint(id), req.Content, userID.(uint))
	if err != nil {
		if errors.Is(err, constants.ErrNotAuthor) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "没有权限更新此评论"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

func (c *CommentController) DeleteCommentGin(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的评论ID"})
		return
	}
	err = c.commentService.DeleteComment(uint(id), userID.(uint))
	if err != nil {
		if errors.Is(err, constants.ErrNotAuthor) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "没有权限删除此评论"})
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{"error": "评论不存在"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
