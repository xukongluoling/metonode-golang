# Personal Blog 2.0

一个基于Golang和Gin框架开发的现代化个人博客系统，具备完整的用户管理、文章管理和评论功能。

## 技术栈

- **后端框架**: Gin - 高性能的Golang Web框架
- **ORM框架**: GORM - 功能强大的Golang ORM库
- **数据库**: MySQL - 可靠的关系型数据库
- **认证机制**: JWT - 无状态的用户认证
- **日志系统**: Zap - 高性能的结构化日志库
- **参数验证**: Validator.v10 - 强大的请求参数验证库

## 功能特性

### 用户管理
- ✅ 用户注册和登录
- ✅ JWT身份认证
- ✅ 密码加密存储
- ✅ 用户信息管理

### 文章管理
- ✅ 文章发布、编辑和删除
- ✅ 文章列表和详情查询
- ✅ 支持富文本内容

### 评论系统
- ✅ 评论发布和删除
- ✅ 按文章查询评论
- ✅ 评论用户信息展示

### 系统功能
- ✅ 统一参数验证
- ✅ 结构化日志记录
- ✅ 错误处理机制
- ✅ 模块化架构设计

## 目录结构

```
personal_blog2/
├── config/            # 配置文件
├── controllers/       # 控制器层
│   └── dto/          # 数据传输对象
├── database/          # 数据库连接
├── middleware/        # 中间件
├── models/            # 数据模型
├── repositories/      # 数据访问层
├── routes/            # 路由配置
├── services/          # 业务逻辑层
├── utils/            # 工具函数
├── main.go           # 入口文件
└── README.md         # 项目说明
```

## 环境要求

- Go 1.18+ 
- MySQL 5.7.0+
- Git

## 快速开始

### 1. 克隆项目

```bash
git clone <项目地址>
cd personal_blog2
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置数据库

编辑 `config/config.yaml` 文件，配置数据库连接信息：

```yaml
server:
  port: "8080"
database:
  driver: "mysql"
  host: "localhost"
  port: "3306"
  username: "root"
  password: "your_password"
  database: "blog_db"
  charset: "utf8mb4"
  parseTime: true
  loc: "Local"
jwt:
  secret: "your_jwt_secret"
  expiration: 7200
```

### 4. 创建数据库

```sql
CREATE DATABASE blog_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 5. 运行项目

```bash
go run main.go
```

项目将在 `http://localhost:8080` 启动。

## API文档

### 用户相关

#### 注册
```
POST /api/v1/users/register
Content-Type: application/json

{
  "username": "testuser",
  "email": "test@example.com",
  "password": "password123"
}
```

#### 登录
```
POST /api/v1/users/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}
```

### 文章相关

#### 创建文章
```
POST /api/v1/posts
Content-Type: application/json
Authorization: Bearer <jwt_token>

{
  "title": "文章标题",
  "content": "文章内容"
}
```

#### 获取文章列表
```
GET /api/v1/posts
```

#### 获取文章详情
```
GET /api/v1/posts/:id
```

### 评论相关

#### 创建评论
```
POST /api/v1/comments
Content-Type: application/json
Authorization: Bearer <jwt_token>

{
  "content": "评论内容",
  "post_id": 1
}
```

#### 获取文章评论
```
GET /api/v1/posts/:post_id/comments
```

## 开发指南

### 代码风格
- 使用Go官方推荐的代码风格
- 变量名采用驼峰命名法
- 函数名采用驼峰命名法，首字母大写表示导出
- 包名使用小写字母

### 日志使用

项目集成了Zap日志库，使用方式：

```go
import (
    "metonode-golang/personal_blog2/utils"
    "go.uber.org/zap"
)

// 记录普通信息
utils.Logger.Info("用户登录成功", zap.String("username", username))

// 记录错误信息
utils.Logger.Error("用户登录失败", zap.String("username", username), zap.Error(err))
```

### 参数验证

使用Validator.v10进行参数验证：

```go
// 在DTO结构体中定义验证规则
type RegisterRequest struct {
    Username string `json:"username" binding:"required,min=3,max=20"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

// 在控制器中使用验证
var request dto.RegisterRequest
if !utils.BindAndValidate(ctx, &request) {
    return
}
```

## 部署说明

### 构建项目

```bash
go build -o personal_blog2 main.go
```

### 运行应用

```bash
./personal_blog2
```

### 环境变量配置

可以通过环境变量覆盖配置文件中的设置：

```bash
# 设置日志级别
export LOG_LEVEL=debug
```

## 日志管理

项目使用Zap日志库，日志格式为JSON，包含以下字段：
- `time`: 日志时间（ISO8601格式）
- `level`: 日志级别（debug/info/warn/error/fatal）
- `msg`: 日志消息
- `caller`: 日志调用位置
- 其他自定义字段（如user_id, request_id等）

### 日志级别配置

通过环境变量 `LOG_LEVEL` 配置日志级别，默认为 `info`：

```bash
export LOG_LEVEL=debug  # 开发环境
export LOG_LEVEL=info   # 生产环境
export LOG_LEVEL=warn   # 线上环境
```
