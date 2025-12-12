package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func main() {
	// 创建登录请求
	loginReq := LoginRequest{
		Username: "testuser",
		Password: "password123",
	}

	// 编码请求体
	jsonData, err := json.Marshal(loginReq)
	if err != nil {
		fmt.Printf("编码请求体失败: %v\n", err)
		return
	}

	// 发送请求
	resp, err := http.Post("http://localhost:8080/api/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("发送请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 打印响应状态码
	fmt.Printf("响应状态码: %d\n", resp.StatusCode)

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应体失败: %v\n", err)
		return
	}

	// 打印响应体内容
	fmt.Printf("响应体内容: %s\n", string(body))

	// 解析响应体
	var loginResp LoginResponse
	if err := json.Unmarshal(body, &loginResp); err != nil {
		fmt.Printf("解析响应体失败: %v\n", err)
		return
	}

	// 打印Token
	fmt.Printf("Token: %s\n", loginResp.Token)
}
