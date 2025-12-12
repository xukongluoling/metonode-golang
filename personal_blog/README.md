# 个人博客 - Go Web 应用

一个现代化、轻量级的个人博客系统，使用 Go 语言开发，采用 RESTful API 设计、JWT 认证和清晰的分层架构。

## 🛠️ 技术栈

- **Web框架**: [Gin](https://github.com/gin-gonic/gin) - 高性能HTTP Web框架
- **ORM**: [GORM](https://gorm.io/) - Go的ORM框架，类似Java的Mybatis
- **数据库**: MySQL - 关系型数据库
- **登录认证**: JWT (JSON Web Tokens) - 无状态认证
- **数据检验**: [validator](https://github.com/go-playground/validator) - Go 结构体校验
- **配置文件**: Viper - 配置管理解决方案
- **密码哈希**: bcrypt - 安全密码加密

## 📁 项目目录结构

```
personal_blog/
├── config/              # 配置管理
│   ├── config.go       # 配置加载逻辑
│   └── config.yaml     # YAML配置文件
├── controllers/        # HTTP请求处理器
│   ├── comment_controller.go
│   ├── post_controller.go
│   └── user_controller.go
├── database/           # 数据库连接和初始化
│   └── db.go
├── global_exceptions/  # 全局错误处理
│   ├── errors.go
│   └── handler_err.go
├── middleware/         # HTTP中间件
│   └── auth.go         # JWT认证中间件
├── models/             # 数据库模型
│   └── models.go
├── repositories/       # 数据访问层
│   ├── comment_repository.go
│   ├── post_repository.go
│   └── user_repository.go
├── routes/             # 路由定义
│   ├── comment_routes.go
│   ├── post_routes.go
│   ├── routes.go
│   └── user_routes.go
├── services/           # 业务逻辑层
│   ├── comment_service.go
│   ├── post_service.go
│   └── user_service.go
├── test/               # 测试文件
│   ├── test_all.go
│   ├── test_login.go
│   └── test_register.go
├── utils/              # 工具函数
│   ├── context.go
│   └── jwt.go
├── main.go             # 应用程序入口
└── README.md           # 项目说明文档
```

## ✨ 功能

- **用户管理**
  - 用户注册
  - 用户登录需要JWT校验
  - 密码加密和检验

- **文章管理**
  - 创建、阅读、更新和删除文章 (CRUD 操作)
  - 获取所有文章或单个文章详情
  - 写操作需要身份认证

- **评论管理**
  - 为文章添加评论
  - 评论操作需要身份认证

- **安全**
  - JWT 令牌认证机制
  - 使用 bcrypt 进行密码哈希
  - 受保护的路由需要身份认证

## 🚀 快速开始

### 前置条件

- Go 1.24 或更高版本
- MySQL 数据库
- Git

### 安装步骤

1. **克隆仓库**
   ```bash
   git clone <仓库地址>
   cd metonode-golang/personal_blog
   ```

2. **安装依赖**
   ```bash
   go mod download
   ```

3. **配置应用程序**
   
   编辑 `config/config.yaml` 文件，配置数据库和 JWT 设置：
   
   ```yaml
   # 服务器配置
   server:
     port: 8080

   # 数据库配置
   mysql:
     host: localhost
     port: 3306
     user: root
     password: your_password
     dbname: personal_blog_task
     charset: utf8mb4
     parseTime: true
     loc: Local

   # JWT 配置
   jwt:
     secret: your_jwt_secret_key
     expireHours: 72
   ```

4. **创建数据库**
   
   ```sql
   CREATE DATABASE personal_blog_task CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
   ```

5. **启动服务器**
   ```bash
   go run main.go
   ```

   服务器将在 `http://localhost:8080` 上启动

### 运行测试

执行测试套件以验证应用程序功能：

```bash
cd test
go run test_all.go
```

## 📝 API 文档

### 认证接口

| 方法 | 端点 | 描述 |
|------|------|------|
| POST | `/blog/api/register` | 注册新用户 |
| POST | `/blog/api/login` | 登录并获取 JWT 令牌 |

### 文章接口

| 方法 | 端点 | 描述 | 是否需要认证 |
|------|------|------|--------------|
| GET | `/blog/api/posts/:id` | 根据 ID 获取文章 | 否 |
| GET | `/blog/api/posts` | 获取所有文章 | 是 |
| POST | `/blog/api/posts/create` | 创建新文章 | 是 |
| PUT | `/blog/api/posts/update/:id` | 更新文章 | 是 |
| DELETE | `/blog/api/posts/delete/:id` | 删除文章 | 是 |

### 评论接口

| 方法 | 端点 | 描述 | 是否需要认证 |
|------|------|------|--------------|
| GET | `/blog/api/comments/:post_id` | 获取文章的所有评论 | 否 |
| POST | `/blog/api/comments/create/:post_id` | 为文章添加评论 | 是 |
| DELETE | `/blog/api/comments/delete/:id` | 删除评论 | 是 |

### 认证方式

受保护的端点需要在 Authorization 头中提供 JWT 令牌：

```
Authorization: Bearer <your-token>
```

## 📊 架构设计

应用程序采用清晰的分层架构模式：

1. **控制器层**: 处理 HTTP 请求和响应
2. **服务层**: 包含业务逻辑
3. **仓库层**: 管理数据访问
4. **模型层**: 定义数据结构
5. **中间件层**: 提供认证等横切关注点

## 🔧 配置说明

应用程序使用 Viper 进行配置管理。配置选项从 `config/config.yaml` 文件加载：

- **Server**: HTTP 服务器端口号
- **MySQL**: 数据库连接详情
- **JWT**: 密钥和令牌过期时间

## 🧪 测试

项目包含完整的测试套件，包括：
- 用户注册测试
- 用户登录测试
- 文章操作测试
- 评论操作测试

运行所有测试：
```bash
cd test
go run test_all.go
```

## 🤝 贡献

欢迎贡献代码！请随时提交 Pull Request。

## 📄 许可证

本项目采用 MIT 许可证 - 查看 LICENSE 文件了解详情。

## 📧 联系方式

如有任何问题或反馈，请联系项目维护者。
