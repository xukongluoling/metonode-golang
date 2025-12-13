package controllers

import (
	"errors"
	"metonode-golang/personal_blog2/constants"
	"metonode-golang/personal_blog2/controllers/dto"
	"metonode-golang/personal_blog2/services"
	"metonode-golang/personal_blog2/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService services.PostService
}

func NewPostController() *PostController {
	return &PostController{postService: services.NewPostService()}
}

func (c *PostController) CreatePost(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req dto.CreatePostRequest
	if !utils.BindAndValidate(ctx, &req) {
		return
	}

	req.UserID = userID.(uint)
	post, err := c.postService.CreatePost(&req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, post)
}

func (c *PostController) GetAllPost(ctx *gin.Context) {
	posts, err := c.postService.GetAllPosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	ctx.JSON(http.StatusOK, posts)
}

func (c *PostController) GetPostByIDGin(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	post, err := c.postService.GetPostByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func (c *PostController) UpdatePost(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}
	var req dto.UpdatePostRequest
	if !utils.BindAndValidate(ctx, &req) {
		return
	}

	req.UserID = userID.(uint)
	req.Id = uint(id)

	post, err := c.postService.UpdatePost(&req)
	if err != nil {
		if errors.Is(err, constants.ErrNotAuthor) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "没有权限更新这篇文章"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func (c *PostController) DeletePost(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}
	if err := c.postService.DeletePost(uint(id), userID.(uint)); err != nil {
		if errors.Is(err, constants.ErrNotAuthor) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "没有权限删除这篇文章"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "文章删除成功"})
}
