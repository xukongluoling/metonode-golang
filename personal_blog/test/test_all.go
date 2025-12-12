package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 定义API基础URL
	baseURL := "http://localhost:8080/blog/api"

	// 1. 测试用户注册
	fmt.Println("1. 测试用户注册...")
	registerData := map[string]string{
		"username": "testuser2",
		"password": "password1235",
		"email":    "test@example.com",
	}
	registerResp, err := sendPostRequest(baseURL+"/register", registerData)
	if err != nil {
		fmt.Printf("注册失败: %v\n", err)
		return
	}
	fmt.Printf("注册响应: %s\n\n", registerResp)

	// 2. 测试用户登录
	fmt.Println("2. 测试用户登录...")
	loginData := map[string]string{
		"username": "testuser2",
		"password": "password1235",
	}
	loginResp, err := sendPostRequest(baseURL+"/login", loginData)
	if err != nil {
		fmt.Printf("登录失败: %v\n", err)
		return
	}
	fmt.Printf("登录响应: %s\n", loginResp)

	// 解析登录响应，获取token
	var loginResult map[string]string
	json.Unmarshal([]byte(loginResp), &loginResult)
	token, exists := loginResult["token"]
	if !exists {
		fmt.Println("未获取到登录token")
		return
	}
	fmt.Printf("获取到token: %s\n\n", token)

	// 3. 测试创建文章
	fmt.Println("3. 测试创建文章...")
	createPostData := map[string]string{
		"title":   "Test Article",
		"content": "This is the content of a test article",
	}
	createPostResp, err := sendPostRequestWithToken(baseURL+"/posts/create", createPostData, token)
	if err != nil {
		fmt.Printf("创建文章失败: %v\n", err)
		return
	}
	fmt.Printf("创建文章响应: %s\n", createPostResp)

	// 解析创建文章响应，获取文章ID
	var postResult map[string]interface{}
	json.Unmarshal([]byte(createPostResp), &postResult)

	// 打印完整响应以便调试
	fmt.Println("创建文章完整响应:", createPostResp)

	// 获取文章ID，尝试两种可能的字段名
	postID, exists := postResult["id"].(float64)
	if !exists {
		// 尝试另一种可能的字段名格式
		postID, exists = postResult["ID"].(float64)
		if !exists {
			fmt.Println("未获取到文章ID")
			return
		}
	}
	fmt.Printf("创建的文章ID: %.0f\n\n", postID)

	// 4. 测试获取所有文章
	fmt.Println("4. 测试获取所有文章...")
	getPostsResp, err := sendGetRequestWithToken(baseURL+"/posts", token)
	if err != nil {
		fmt.Printf("获取所有文章失败: %v\n", err)
		return
	}
	fmt.Printf("获取所有文章响应: %s\n\n", getPostsResp)

	// 5. 测试获取文章详情
	fmt.Printf("5. 测试获取文章详情（ID: %.0f）...\n", postID)
	getPostResp, err := sendGetRequestWithToken(fmt.Sprintf("%s/posts/%.0f", baseURL, postID), token)
	if err != nil {
		fmt.Printf("获取文章详情失败: %v\n", err)
		return
	}
	fmt.Printf("获取文章详情响应: %s\n\n", getPostResp)

	// 6. 测试创建评论
	fmt.Println("6. 测试创建评论...")
	createCommentData := map[string]interface{}{
		"content": "This is a test comment",
		"post_id": postID,
	}
	createCommentResp, err := sendPostRequestWithToken(baseURL+"/comments/create", createCommentData, token)
	if err != nil {
		fmt.Printf("创建评论失败: %v\n", err)
		return
	}
	fmt.Printf("创建评论响应: %s\n", createCommentResp)

	// 解析创建评论响应，获取评论ID
	var commentResult map[string]interface{}
	json.Unmarshal([]byte(createCommentResp), &commentResult)

	// 打印完整响应以便调试
	fmt.Println("创建评论完整响应:", createCommentResp)

	// 获取评论ID，尝试两种可能的字段名
	commentID, exists := commentResult["id"].(float64)
	if !exists {
		// 尝试另一种可能的字段名格式
		commentID, exists = commentResult["ID"].(float64)
		if !exists {
			fmt.Println("未获取到评论ID")
			return
		}
	}
	fmt.Printf("创建的评论ID: %.0f\n\n", commentID)

	// 7. 测试获取文章评论
	fmt.Printf("7. 测试获取文章评论（文章ID: %.0f）...\n", postID)
	getCommentsResp, err := sendGetRequestWithToken(fmt.Sprintf("%s/comments/%.0f", baseURL, postID), token)
	if err != nil {
		fmt.Printf("获取文章评论失败: %v\n", err)
		return
	}
	fmt.Printf("获取文章评论响应: %s\n\n", getCommentsResp)

	// 8. 测试更新文章
	fmt.Printf("8. 测试更新文章（ID: %.0f）...\n", postID)
	updatePostData := map[string]string{
		"title":   "更新后的测试文章",
		"content": "这是更新后的测试文章内容",
	}
	updatePostResp, err := sendPutRequestWithToken(fmt.Sprintf("%s/posts/update/%.0f", baseURL, postID), updatePostData, token)
	if err != nil {
		fmt.Printf("更新文章失败: %v\n", err)
		return
	}
	fmt.Printf("更新文章响应: %s\n\n", updatePostResp)

	// 9. 测试更新评论
	fmt.Printf("9. 测试更新评论（ID: %.0f）...\n", commentID)
	updateCommentData := map[string]string{
		"content": "这是更新后的测试评论",
	}
	updateCommentResp, err := sendPutRequestWithToken(fmt.Sprintf("%s/comments/update/%.0f", baseURL, commentID), updateCommentData, token)
	if err != nil {
		fmt.Printf("更新评论失败: %v\n", err)
		return
	}
	fmt.Printf("更新评论响应: %s\n\n", updateCommentResp)

	// 10. 测试删除评论
	fmt.Printf("10. 测试删除评论（ID: %.0f）...\n", commentID)
	deleteCommentResp, err := sendDeleteRequestWithToken(fmt.Sprintf("%s/comments/delete/%.0f", baseURL, commentID), token)
	if err != nil {
		fmt.Printf("删除评论失败: %v\n", err)
		return
	}
	fmt.Printf("删除评论响应: %s\n\n", deleteCommentResp)

	// 11. 测试删除文章
	fmt.Printf("11. 测试删除文章（ID: %.0f）...\n", postID)
	deletePostResp, err := sendDeleteRequestWithToken(fmt.Sprintf("%s/posts/delete/%.0f", baseURL, postID), token)
	if err != nil {
		fmt.Printf("删除文章失败: %v\n", err)
		return
	}
	fmt.Printf("删除文章响应: %s\n\n", deletePostResp)

	fmt.Println("所有测试完成！")
}

// sendPostRequest 发送POST请求
func sendPostRequest(url string, data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// sendPostRequestWithToken 发送带token的POST请求
func sendPostRequestWithToken(url string, data interface{}, token string) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// sendPutRequestWithToken 发送带token的PUT请求
func sendPutRequestWithToken(url string, data interface{}, token string) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// sendDeleteRequestWithToken 发送带token的DELETE请求
func sendDeleteRequestWithToken(url string, token string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// sendGetRequestWithToken 发送带token的GET请求
func sendGetRequestWithToken(url string, token string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
