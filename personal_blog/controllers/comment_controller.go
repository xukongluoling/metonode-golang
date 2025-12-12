package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"metonode-golang/personal_blog/global_exceptions"
	"metonode-golang/personal_blog/models"
	"metonode-golang/personal_blog/services"
	"metonode-golang/personal_blog/utils"

	"github.com/gin-gonic/gin"
)

// CommentController 评论控制器
type CommentController struct {
	commentService *services.CommentService
}

// NewCommentController 创建评论控制器实例
func NewCommentController() *CommentController {
	return &CommentController{
		commentService: services.NewCommentService(),
	}
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

// CreateComment 创建评论 (兼容标准库)
func (c *CommentController) CreateComment(w http.ResponseWriter, r *http.Request) {
	// 从上下文获取用户ID
	userID, ok := utils.GetUserIDFromContext(r.Context())
	if !ok {
		global_exceptions.HandlerError(w, global_exceptions.ErrUnauthorized)
		return
	}

	var req CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		global_exceptions.HandlerError(w, global_exceptions.ErrBadRequest)
		return
	}

	comment, err := c.commentService.CreateComment(req.Content, req.PostID, userID)
	if err != nil {
		global_exceptions.HandlerError(w, global_exceptions.NewAppError(http.StatusBadRequest, err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

// CreateCommentGin 创建评论 (Gin版本)
func (c *CommentController) CreateCommentGin(ctx *gin.Context) {
	// 从Gin上下文获取用户ID
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	var req CreateCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	comment, err := c.commentService.CreateComment(req.Content, req.PostID, userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

// GetCommentsByPostID 根据文章ID获取所有评论 (兼容标准库)
func (c *CommentController) GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	// 解析文章ID
	idStr := r.URL.Path[len("/api/posts/"):]
	idStr = idStr[:len(idStr)-len("/comments")]
	postID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		global_exceptions.HandlerError(w, global_exceptions.ErrBadRequest)
		return
	}

	comments, err := c.commentService.GetCommentsByPostID(uint(postID))
	if err != nil {
		global_exceptions.HandlerError(w, global_exceptions.ErrNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}

// GetCommentsByPostIDGin 根据文章ID获取所有评论 (Gin版本)
func (c *CommentController) GetCommentsByPostIDGin(ctx *gin.Context) {
	// 从Gin上下文获取路径参数
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

// UpdateComment 更新评论 (兼容标准库)
func (c *CommentController) UpdateComment(w http.ResponseWriter, r *http.Request) {
	// 从上下文获取用户ID
	userID, ok := utils.GetUserIDFromContext(r.Context())
	if !ok {
		global_exceptions.HandlerError(w, global_exceptions.ErrUnauthorized)
		return
	}

	// 解析评论ID
	idStr := r.URL.Path[len("/api/comments/update/"):]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		global_exceptions.HandlerError(w, global_exceptions.ErrBadRequest)
		return
	}

	var req UpdateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		global_exceptions.HandlerError(w, global_exceptions.ErrBadRequest)
		return
	}

	comment, err := c.commentService.UpdateComment(uint(id), req.Content, userID)
	if err != nil {
		if errors.Is(err, models.ErrNotAuthor) {
			global_exceptions.HandlerError(w, global_exceptions.ErrForbidden)
			return
		}
		global_exceptions.HandlerError(w, global_exceptions.NewAppError(http.StatusBadRequest, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comment)
}

// UpdateCommentGin 更新评论 (Gin版本)
func (c *CommentController) UpdateCommentGin(ctx *gin.Context) {
	// 从Gin上下文获取用户ID
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	// 从Gin上下文获取路径参数
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的评论ID"})
		return
	}

	var req UpdateCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	comment, err := c.commentService.UpdateComment(uint(id), req.Content, userID.(uint))
	if err != nil {
		if errors.Is(err, models.ErrNotAuthor) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "没有权限更新此评论"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// DeleteComment 删除评论 (兼容标准库)
func (c *CommentController) DeleteComment(w http.ResponseWriter, r *http.Request) {
	// 从上下文获取用户ID
	userID, ok := utils.GetUserIDFromContext(r.Context())
	if !ok {
		global_exceptions.HandlerError(w, global_exceptions.ErrUnauthorized)
		return
	}

	// 解析评论ID
	idStr := r.URL.Path[len("/api/comments/delete/"):]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		global_exceptions.HandlerError(w, global_exceptions.ErrBadRequest)
		return
	}

	if err := c.commentService.DeleteComment(uint(id), userID); err != nil {
		if errors.Is(err, models.ErrNotAuthor) {
			global_exceptions.HandlerError(w, global_exceptions.ErrForbidden)
			return
		}
		global_exceptions.HandlerError(w, global_exceptions.ErrNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "评论删除成功"})
}

// DeleteCommentGin 删除评论 (Gin版本)
func (c *CommentController) DeleteCommentGin(ctx *gin.Context) {
	// 从Gin上下文获取用户ID
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	// 从Gin上下文获取路径参数
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的评论ID"})
		return
	}

	if err := c.commentService.DeleteComment(uint(id), userID.(uint)); err != nil {
		if errors.Is(err, models.ErrNotAuthor) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "没有权限删除此评论"})
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{"error": "评论不存在"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "评论删除成功"})
}
