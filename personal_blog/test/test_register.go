package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type RegisterResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	// 创建注册请求
	registerReq := RegisterRequest{
		Username: "testuser",
		Password: "password123",
		Nickname: "测试用户",
	}

	// 编码请求体
	jsonData, err := json.Marshal(registerReq)
	if err != nil {
		fmt.Printf("编码请求体失败: %v\n", err)
		return
	}

	// 发送请求
	resp, err := http.Post("http://localhost:8080/api/register", "application/json", bytes.NewBuffer(jsonData))
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
	var registerResp RegisterResponse
	if err := json.Unmarshal(body, &registerResp); err != nil {
		fmt.Printf("解析响应体失败: %v\n", err)
		return
	}

	// 打印响应信息
	fmt.Printf("注册响应信息: %s\n", registerResp.Message)
}
