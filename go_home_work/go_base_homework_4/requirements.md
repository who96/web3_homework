# 个人博客后端系统 - Requirements Document

使用Go+Gin+GORM+MySQL开发个人博客系统，包括用户认证、文章CRUD、评论功能，提供完整API文档和自动化测试

## Core Features

个人博客系统的核心数据流：用户 -> 认证 -> 文章操作 -> 评论交互

**核心实体:**
- User: 用户身份管理
- Post: 博客文章内容 
- Comment: 文章评论交互
- Authentication: JWT令牌认证

**核心功能:**
1. 用户认证系统：注册、登录、JWT验证
2. 文章管理系统：创建、读取、更新、删除 (CRUD)
3. 评论交互系统：评论发布、评论列表
4. 权限控制系统：作者权限、访客权限

## User Stories

**用户身份管理**
- As a 访客, I want 注册账户, so that 我能创建和管理自己的博客文章
- As a 用户, I want 安全登录, so that 我能访问需要权限的功能
- As a 用户, I want JWT持久化认证, so that 我不用重复登录

**文章管理**  
- As a 认证用户, I want 创建博客文章, so that 我能分享自己的想法
- As a 任何人, I want 浏览文章列表, so that 我能发现感兴趣的内容
- As a 任何人, I want 阅读单篇文章, so that 我能获取完整信息
- As a 文章作者, I want 更新自己的文章, so that 我能修正或补充内容
- As a 文章作者, I want 删除自己的文章, so that 我能管理个人内容

**评论交互**
- As a 认证用户, I want 对文章发表评论, so that 我能与作者互动
- As a 任何人, I want 查看文章评论, so that 我能了解他人观点

## Acceptance Criteria

**项目基础**
- [x] Go项目初始化，使用go mod管理依赖
- [x] 集成Gin框架、GORM库、MySQL驱动
- [x] 数据库连接配置：mysql root:root@localhost:3306/web3_homework

**数据库设计**
- [x] 创建users表：id, username, password, email, created_at, updated_at
- [x] 创建posts表：id, title, content, user_id, created_at, updated_at  
- [x] 创建comments表：id, content, user_id, post_id, created_at
- [x] GORM模型定义并自动迁移

**用户认证功能**
- [x] 用户注册：验证输入，密码bcrypt加密，存储用户信息
- [x] 测试用户注册：curl测试正常注册和异常情况（用户名重复、参数缺失等）
- [x] 用户登录：验证用户名密码，生成JWT返回
- [x] 测试用户登录：curl测试正确登录和错误密码、用户不存在等情况
- [x] JWT认证中间件：验证token有效性，提取用户信息
- [x] 测试JWT中间件：curl测试有效token、无效token、过期token等情况

**文章管理功能**
- [x] 创建文章：认证用户提供标题内容，保存到数据库
- [x] 测试创建文章：curl测试有token创建、无token创建、参数缺失等情况
- [x] 文章列表：返回所有文章基本信息，包含作者信息
- [x] 测试文章列表：curl测试获取文章列表，验证返回数据格式
- [x] 文章详情：根据ID返回完整文章内容
- [x] 测试文章详情：curl测试正确ID、错误ID、不存在ID等情况
- [x] 更新文章：验证作者权限，更新标题内容
- [x] 测试更新文章：curl测试作者更新、非作者更新、文章不存在等情况
- [x] 删除文章：验证作者权限，从数据库删除
- [x] 测试删除文章：curl测试作者删除、非作者删除、文章不存在等情况

**评论功能** 
- [x] 创建评论：认证用户对指定文章发表评论
- [x] 测试创建评论：curl测试有token评论、无token评论、文章不存在等情况
- [x] 评论列表：返回指定文章的所有评论，包含评论者信息
- [x] 测试评论列表：curl测试获取评论列表，验证返回数据格式

**文档与测试**
- [x] README文档：环境配置、依赖安装、启动步骤
- [x] Postman接口文档：所有API的完整请求响应示例
- [x] 自动化测试脚本：一键测试所有接口功能

## Non-functional Requirements

**开发流程要求**
- 接口驱动开发：每完成一个接口立即进行功能测试
- 增量验证策略：确保当前接口完全可用后再开发下一个
- 问题隔离原则：避免问题累积，及时发现和修复缺陷
- 单接口测试：使用curl或Postman验证每个接口的正常和异常流程

**错误处理**
- 统一错误响应格式，返回适当HTTP状态码
- 数据库错误处理，连接失败、约束违反等
- 认证失败处理，token无效、权限不足等
- 资源不存在处理，文章、用户、评论不存在等

**性能要求**
- 数据库查询优化：使用GORM预加载减少N+1查询
- 响应时间：单个API请求响应时间 < 100ms
- 并发支持：支持100个并发用户同时访问

**安全要求** 
- 密码安全：使用bcrypt哈希算法，成本因子≥10
- JWT安全：使用HS256算法，密钥长度≥32位，过期时间24小时
- 输入验证：所有用户输入进行参数验证，防止SQL注入
- 权限控制：严格验证用户只能操作自己的资源

**兼容要求**
- Go版本：≥1.19
- MySQL版本：≥5.7，运行在Docker容器中
- API规范：RESTful风格，JSON格式数据交换
- HTTP状态码：遵循标准HTTP状态码语义
