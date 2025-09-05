# 个人博客后端系统

使用Go+Gin+GORM+MySQL开发的个人博客系统，支持用户认证、文章管理和评论功能。

## 项目概述

这是一个完整的个人博客后端API系统，提供以下功能：

- **用户认证**：用户注册、登录、JWT认证
- **文章管理**：创建、查看、更新、删除文章（CRUD操作）
- **评论系统**：对文章发表评论、查看评论列表
- **权限控制**：基于JWT的用户认证和资源权限验证

## 技术栈

- **Go 1.23+**：主要编程语言
- **Gin**：Web框架
- **GORM**：ORM库
- **MySQL**：数据库（运行在Docker中）
- **JWT**：用户认证
- **bcrypt**：密码加密

## 环境要求

- Go 1.23或更高版本
- MySQL 5.7+（推荐使用Docker运行）
- Docker（可选，用于运行MySQL）

## 安装和启动

### 1. 克隆项目

```bash
cd go_base_homework_4
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 启动MySQL数据库

确保MySQL运行在本地Docker容器中：

```bash
# 检查MySQL容器是否运行
docker ps | grep mysql

# 如果没有运行，启动MySQL容器
docker run -d --name mysql-blog \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=root \
  -e MYSQL_DATABASE=web3_homework \
  mysql:latest
```

### 4. 启动服务器

```bash
go run .
```

服务器将在 http://localhost:8080 启动

## API接口文档

### 用户认证

#### 注册用户
- **POST** `/api/auth/register`
- **请求体**:
```json
{
  "username": "testuser",
  "password": "password123", 
  "email": "test@example.com"
}
```

#### 用户登录
- **POST** `/api/auth/login`
- **请求体**:
```json
{
  "username": "testuser",
  "password": "password123"
}
```

### 文章管理

#### 获取文章列表
- **GET** `/api/posts`

#### 获取文章详情
- **GET** `/api/posts/{id}`

#### 创建文章（需要认证）
- **POST** `/api/protected/posts`
- **Headers**: `Authorization: Bearer {jwt_token}`
- **请求体**:
```json
{
  "title": "文章标题",
  "content": "文章内容"
}
```

#### 更新文章（需要认证，仅作者）
- **PUT** `/api/protected/posts/{id}`
- **Headers**: `Authorization: Bearer {jwt_token}`
- **请求体**:
```json
{
  "title": "更新的标题",
  "content": "更新的内容"
}
```

#### 删除文章（需要认证，仅作者）
- **DELETE** `/api/protected/posts/{id}`
- **Headers**: `Authorization: Bearer {jwt_token}`

### 评论管理

#### 获取文章评论列表
- **GET** `/api/posts/{id}/comments`

#### 创建评论（需要认证）
- **POST** `/api/protected/posts/{id}/comments`
- **Headers**: `Authorization: Bearer {jwt_token}`
- **请求体**:
```json
{
  "content": "这是一条评论"
}
```

## 测试

### 快速测试流程

1. **用户注册**
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "password123", "email": "test@example.com"}'
```

2. **用户登录获取token**
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "password123"}'
```

3. **创建文章**
```bash
curl -X POST http://localhost:8080/api/protected/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"title": "测试文章", "content": "这是一篇测试文章"}'
```

4. **查看文章列表**
```bash
curl http://localhost:8080/api/posts
```

## 数据库结构

### users表
- `id`: 用户ID (主键)
- `username`: 用户名 (唯一)
- `password`: 加密密码
- `email`: 邮箱 (唯一)
- `created_at`: 创建时间
- `updated_at`: 更新时间

### posts表
- `id`: 文章ID (主键)
- `title`: 文章标题
- `content`: 文章内容
- `user_id`: 作者ID (外键)
- `created_at`: 创建时间
- `updated_at`: 更新时间

### comments表
- `id`: 评论ID (主键)
- `content`: 评论内容
- `user_id`: 评论者ID (外键)
- `post_id`: 文章ID (外键)
- `created_at`: 创建时间

## 安全特性

- 密码使用bcrypt加密存储
- JWT令牌认证，有效期24小时
- 权限控制：只有作者能修改/删除自己的文章
- 输入验证和错误处理
- SQL注入防护（GORM内置）

## 开发者信息

本项目按照Linus Torvalds的"好品味"原则开发：
- 简洁明了的数据结构设计
- 消除不必要的复杂性
- 实用主义优先
- 严格的错误处理

所有接口都经过严格测试，确保在正常和异常情况下都能正确响应。