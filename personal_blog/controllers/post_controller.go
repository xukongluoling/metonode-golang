package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"metonode-golang/personal_blog/global_exceptions"
	"metonode-golang/personal_blog/models"
	"metonode-golang/personal_blog/services"
	"metonode-golang/personal_blog/utils"

	"github.com/gin-gonic/gin"
)

// PostController 文章控制器
type PostController struct {
	postService *services.PostService
}

// NewPostController 创建文章控制器实例
func NewPostController() *PostController {
	return &PostController{
		postService: services.NewPostService(),
	}
}

// CreatePostRequest 创建文章请求结构体
type CreatePostRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

// UpdatePostRequest 更新文章请求结构体
type UpdatePostRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

// CreatePost 创建文章 (兼容标准库)
func (c *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	// 从上下文获取用户ID
	userID, ok := utils.GetUserIDFromContext(r.Context())
	if !ok {
		global_exceptions.HandlerError(w, global_exceptions.ErrUnauthorized)
		return
	}

	var req CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		global_exceptions.HandlerError(w, global_exceptions.ErrBadRequest)
		return
	}

	post, err := c.postService.CreatePost(req.Title, req.Content, userID)
	if err != nil {
		global_exceptions.HandlerError(w, global_exceptions.NewAppError(http.StatusBadRequest, err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

// CreatePostGin 创建文章 (Gin版本)
func (c *PostController) CreatePostGin(ctx *gin.Context) {
	// 从Gin上下文获取用户ID
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	var req CreatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	post, err := c.postService.CreatePost(req.Title, req.Content, userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, post)
}

// GetAllPosts 获取所有文章 (兼容标准库)
func (c *PostController) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := c.postService.GetAllPosts()
	if err != nil {
		global_exceptions.HandlerError(w, global_exceptions.ErrInternalServer)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

// GetAllPostsGin 获取所有文章 (Gin版本)
func (c *PostController) GetAllPostsGin(ctx *gin.Context) {
	posts, err := c.postService.GetAllPosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "服务器内部错误"})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

// GetPostByID 根据ID获取文章 (兼容标准库)
func (c *PostController) GetPostByID(w http.ResponseWriter, r *http.Request) {
	// 解析文章ID
	idStr := r.URL.Path[len("/api/posts/"):]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		global_exceptions.HandlerError(w, global_exceptions.ErrBadRequest)
		return
	}

	post, err := c.postService.GetPostByID(uint(id))
	if err != nil {
		global_exceptions.HandlerError(w, global_exceptions.ErrNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

// GetPostByIDGin 根据ID获取文章 (Gin版本)
func (c *PostController) GetPostByIDGin(ctx *gin.Context) {
	// 从Gin上下文获取路径参数
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	post, err := c.postService.GetPostByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

// UpdatePost 更新文章 (兼容标准库)
func (c *PostController) UpdatePost(w http.ResponseWriter, r *http.Request) {
	// 从上下文获取用户ID
	userID, ok := utils.GetUserIDFromContext(r.Context())
	if !ok {
		global_exceptions.HandlerError(w, global_exceptions.ErrUnauthorized)
		return
	}

	// 解析文章ID
	idStr := r.URL.Path[len("/api/posts/update/"):]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		global_exceptions.HandlerError(w, global_exceptions.ErrBadRequest)
		return
	}

	var req UpdatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		global_exceptions.HandlerError(w, global_exceptions.ErrBadRequest)
		return
	}

	post, err := c.postService.UpdatePost(uint(id), req.Title, req.Content, userID)
	if err != nil {
		if err == models.ErrNotAuthor {
			global_exceptions.HandlerError(w, global_exceptions.ErrForbidden)
			return
		}
		global_exceptions.HandlerError(w, global_exceptions.NewAppError(http.StatusBadRequest, err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

// UpdatePostGin 更新文章 (Gin版本)
func (c *PostController) UpdatePostGin(ctx *gin.Context) {
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	var req UpdatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	post, err := c.postService.UpdatePost(uint(id), req.Title, req.Content, userID.(uint))
	if err != nil {
		if err == models.ErrNotAuthor {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "没有权限更新此文章"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

// DeletePost 删除文章 (兼容标准库)
func (c *PostController) DeletePost(w http.ResponseWriter, r *http.Request) {
	// 从上下文获取用户ID
	userID, ok := utils.GetUserIDFromContext(r.Context())
	if !ok {
		global_exceptions.HandlerError(w, global_exceptions.ErrUnauthorized)
		return
	}

	// 解析文章ID
	idStr := r.URL.Path[len("/api/posts/delete/"):]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		global_exceptions.HandlerError(w, global_exceptions.ErrBadRequest)
		return
	}

	if err := c.postService.DeletePost(uint(id), userID); err != nil {
		if err == models.ErrNotAuthor {
			global_exceptions.HandlerError(w, global_exceptions.ErrForbidden)
			return
		}
		global_exceptions.HandlerError(w, global_exceptions.ErrNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "文章删除成功"})
}

// DeletePostGin 删除文章 (Gin版本)
func (c *PostController) DeletePostGin(ctx *gin.Context) {
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	if err := c.postService.DeletePost(uint(id), userID.(uint)); err != nil {
		if err == models.ErrNotAuthor {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "没有权限删除此文章"})
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "文章删除成功"})
}
