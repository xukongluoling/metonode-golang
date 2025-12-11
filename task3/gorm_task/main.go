package main

import (
	"flag"
	"log"
	"metonode-golang/task3/gorm_task/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 定义命令行参数，默认查询ID为1的用户
	userId := flag.Uint64("userId", 1, "要查询的用户ID")
	flag.Parse()

	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, er := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if er != nil {
		log.Fatal("mysql connect err:", er)
		return
	}
	log.Println("mysql connect success")
	// Gorm创建或更新这些模型对应的数据库表。
	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}); err != nil {
		log.Fatal("AutoMigrate err:", err)
	}
	log.Println("AutoMigrate success")

	// 测试钩子函数功能
	log.Println("=== 开始测试钩子函数功能 ===")

	// 创建测试用户
	var testUser models.User
	testUser.Name = "test_user"
	testUser.Password = "password123"
	if err := db.FirstOrCreate(&testUser, models.User{Name: testUser.Name}).Error; err != nil {
		log.Printf("创建测试用户失败: %v", err)
	} else {
		log.Printf("创建/获取测试用户成功: ID=%d, 用户名=%s, 文章数量=%d", testUser.Id, testUser.Name, testUser.PostCount)
	}

	// 创建测试文章
	var testPost models.Post
	testPost.Title = "测试文章标题"
	testPost.Content = "测试文章内容"
	testPost.UserID = testUser.Id
	if err := db.Create(&testPost).Error; err != nil {
		log.Printf("创建测试文章失败: %v", err)
	} else {
		// 检查用户文章数量是否更新
		var updatedUser models.User
		if err := db.First(&updatedUser, testUser.Id).Error; err != nil {
			log.Printf("查询更新后的用户信息失败: %v", err)
		} else {
			log.Printf("创建文章后，用户信息更新: ID=%d, 用户名=%s, 文章数量=%d", updatedUser.Id, updatedUser.Name, updatedUser.PostCount)
		}

		// 检查文章的评论状态
		var updatedPost models.Post
		if err := db.First(&updatedPost, testPost.Id).Error; err != nil {
			log.Printf("查询更新后的文章信息失败: %v", err)
		} else {
			log.Printf("创建文章后，文章信息: ID=%d, 标题=%s, 评论状态=%s", updatedPost.Id, updatedPost.Title, updatedPost.CommentStatus)
		}
	}

	// 创建测试评论
	var testComment models.Comment
	testComment.PostId = testPost.Id
	testComment.UserId = testUser.Id
	testComment.Username = testUser.Name
	testComment.Content = "测试评论内容"
	if err := db.Create(&testComment).Error; err != nil {
		log.Printf("创建测试评论失败: %v", err)
	} else {
		// 检查文章的评论状态
		var updatedPost models.Post
		if err := db.First(&updatedPost, testPost.Id).Error; err != nil {
			log.Printf("查询更新后的文章信息失败: %v", err)
		} else {
			log.Printf("创建评论后，文章信息: ID=%d, 标题=%s, 评论状态=%s", updatedPost.Id, updatedPost.Title, updatedPost.CommentStatus)
		}
	}

	// 删除测试评论
	if err := db.Delete(&testComment).Error; err != nil {
		log.Printf("删除测试评论失败: %v", err)
	} else {
		// 检查文章的评论状态
		var updatedPost models.Post
		if err := db.First(&updatedPost, testPost.Id).Error; err != nil {
			log.Printf("查询更新后的文章信息失败: %v", err)
		} else {
			log.Printf("删除评论后，文章信息: ID=%d, 标题=%s, 评论状态=%s", updatedPost.Id, updatedPost.Title, updatedPost.CommentStatus)
		}
	}

	log.Println("=== 钩子函数功能测试完成 ===")

	// 1. 查询指定用户ID的所有文章及其对应的评论信息
	var user models.User
	if err := db.Preload("Posts.Comments").First(&user, *userId).Error; err != nil {
		log.Printf("查询用户失败: %v", err)
	} else {
		log.Printf("用户 %s 的文章及其评论:", user.Name)
		for _, post := range user.Posts {
			log.Printf("  文章: %s (ID: %d)", post.Title, post.Id)
			log.Printf("  评论数: %d", len(post.Comments))
			for _, comment := range post.Comments {
				log.Printf("    评论者: %s, 内容: %s", comment.Username, comment.Content)
			}
		}
	}

	// 2. 查询评论数量最多的文章信息
	var mostCommentedPost models.Post
	if err := db.Preload("Comments").Preload("User").Order("(SELECT COUNT(*) FROM comments WHERE comments.post_id = posts.id) DESC").First(&mostCommentedPost).Error; err != nil {
		log.Printf("查询评论最多的文章失败: %v", err)
	} else {
		log.Printf("评论数量最多的文章:")
		log.Printf("  标题: %s", mostCommentedPost.Title)
		log.Printf("  作者: %s", mostCommentedPost.User.Name)
		log.Printf("  内容: %s", mostCommentedPost.Content)
		log.Printf("  评论数: %d", len(mostCommentedPost.Comments))
		log.Printf("  评论列表:")
		for _, comment := range mostCommentedPost.Comments {
			log.Printf("    评论者: %s, 内容: %s", comment.Username, comment.Content)
		}
	}
}
