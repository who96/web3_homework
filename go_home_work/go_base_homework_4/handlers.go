package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RegisterRequest 注册请求结构
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6,max=100"`
	Email    string `json:"email" binding:"required,email"`
}

// RegisterResponse 注册响应结构
type RegisterResponse struct {
	Message string `json:"message"`
	UserID  uint   `json:"user_id,omitempty"`
}

// ErrorResponse 错误响应结构
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// Register 用户注册处理器
func Register(c *gin.Context) {
	var req RegisterRequest
	
	// 绑定并验证JSON请求
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: "请求参数验证失败: " + err.Error(),
		})
		return
	}
	
	// 检查用户名是否已存在
	var existingUser User
	if err := db.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, ErrorResponse{
			Error:   "username_exists",
			Message: "用户名已存在",
		})
		return
	}
	
	// 检查邮箱是否已存在
	if err := db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, ErrorResponse{
			Error:   "email_exists",
			Message: "邮箱已被注册",
		})
		return
	}
	
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "password_hash_error",
			Message: "密码加密失败",
		})
		return
	}
	
	// 创建新用户
	user := User{
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
	}
	
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "database_error",
			Message: "创建用户失败: " + err.Error(),
		})
		return
	}
	
	// 成功响应
	c.JSON(http.StatusCreated, RegisterResponse{
		Message: "用户注册成功",
		UserID:  user.ID,
	})
}

// LoginRequest 登录请求结构
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应结构
type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	UserID  uint   `json:"user_id"`
}

// Login 用户登录处理器
func Login(c *gin.Context) {
	var req LoginRequest
	
	// 绑定并验证JSON请求
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: "请求参数验证失败: " + err.Error(),
		})
		return
	}
	
	// 查找用户
	var user User
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   "invalid_credentials",
			Message: "用户名或密码错误",
		})
		return
	}
	
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   "invalid_credentials", 
			Message: "用户名或密码错误",
		})
		return
	}
	
	// 生成JWT
	token, err := GenerateJWT(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "token_generation_error",
			Message: "生成访问令牌失败",
		})
		return
	}
	
	// 成功响应
	c.JSON(http.StatusOK, LoginResponse{
		Message: "登录成功",
		Token:   token,
		UserID:  user.ID,
	})
}

// CreatePostRequest 创建文章请求结构
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=200"`
	Content string `json:"content" binding:"required,min=1"`
}

// PostResponse 文章响应结构
type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	Username  string    `json:"username"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

// CreatePost 创建文章处理器
func CreatePost(c *gin.Context) {
	var req CreatePostRequest
	
	// 绑定并验证JSON请求
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: "请求参数验证失败: " + err.Error(),
		})
		return
	}
	
	// 获取当前用户ID
	userID := GetCurrentUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   "unauthorized",
			Message: "无法获取用户信息",
		})
		return
	}
	
	// 创建文章
	post := Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}
	
	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "database_error",
			Message: "创建文章失败: " + err.Error(),
		})
		return
	}
	
	// 获取用户信息用于响应
	var user User
	db.First(&user, userID)
	
	// 成功响应
	c.JSON(http.StatusCreated, PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		Username:  user.Username,
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}

// GetPosts 获取文章列表处理器
func GetPosts(c *gin.Context) {
	var posts []Post
	
	// 预加载用户信息，避免N+1查询问题
	if err := db.Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "database_error",
			Message: "获取文章列表失败: " + err.Error(),
		})
		return
	}
	
	// 转换为响应格式
	var postResponses []PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, PostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			UserID:    post.UserID,
			Username:  post.User.Username,
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	
	// 成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "获取文章列表成功",
		"data":    postResponses,
		"count":   len(postResponses),
	})
}

// GetPostByID 获取文章详情处理器
func GetPostByID(c *gin.Context) {
	postID := c.Param("id")
	
	var post Post
	// 预加载用户信息
	if err := db.Preload("User").First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "post_not_found",
			Message: "文章不存在",
		})
		return
	}
	
	// 转换为响应格式
	postResponse := PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		Username:  post.User.Username,
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	
	// 成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "获取文章详情成功",
		"data":    postResponse,
	})
}

// UpdatePostRequest 更新文章请求结构
type UpdatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=200"`
	Content string `json:"content" binding:"required,min=1"`
}

// UpdatePost 更新文章处理器
func UpdatePost(c *gin.Context) {
	postID := c.Param("id")
	var req UpdatePostRequest
	
	// 绑定并验证JSON请求
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: "请求参数验证失败: " + err.Error(),
		})
		return
	}
	
	// 获取当前用户ID
	userID := GetCurrentUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   "unauthorized",
			Message: "无法获取用户信息",
		})
		return
	}
	
	// 查找文章并验证作者权限
	var post Post
	if err := db.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "post_not_found",
			Message: "文章不存在",
		})
		return
	}
	
	// 验证是否为文章作者
	if post.UserID != userID {
		c.JSON(http.StatusForbidden, ErrorResponse{
			Error:   "access_denied",
			Message: "只有文章作者才能修改文章",
		})
		return
	}
	
	// 更新文章
	post.Title = req.Title
	post.Content = req.Content
	
	if err := db.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "database_error",
			Message: "更新文章失败: " + err.Error(),
		})
		return
	}
	
	// 获取用户信息用于响应
	var user User
	db.First(&user, userID)
	
	// 成功响应
	c.JSON(http.StatusOK, PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		Username:  user.Username,
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}

// DeletePost 删除文章处理器
func DeletePost(c *gin.Context) {
	postID := c.Param("id")
	
	// 获取当前用户ID
	userID := GetCurrentUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   "unauthorized",
			Message: "无法获取用户信息",
		})
		return
	}
	
	// 查找文章并验证作者权限
	var post Post
	if err := db.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "post_not_found",
			Message: "文章不存在",
		})
		return
	}
	
	// 验证是否为文章作者
	if post.UserID != userID {
		c.JSON(http.StatusForbidden, ErrorResponse{
			Error:   "access_denied",
			Message: "只有文章作者才能删除文章",
		})
		return
	}
	
	// 先删除文章相关评论
	if err := db.Where("post_id = ?", post.ID).Delete(&Comment{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "database_error",
			Message: "删除文章评论失败: " + err.Error(),
		})
		return
	}
	
	// 删除文章
	if err := db.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "database_error",
			Message: "删除文章失败: " + err.Error(),
		})
		return
	}
	
	// 成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "文章删除成功",
		"post_id": post.ID,
	})
}

// CreateCommentRequest 创建评论请求结构
type CreateCommentRequest struct {
	Content string `json:"content" binding:"required,min=1,max=1000"`
}

// CommentResponse 评论响应结构
type CommentResponse struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	PostID    uint   `json:"post_id"`
	CreatedAt string `json:"created_at"`
}

// CreateComment 创建评论处理器
func CreateComment(c *gin.Context) {
	postID := c.Param("id")
	var req CreateCommentRequest
	
	// 绑定并验证JSON请求
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: "请求参数验证失败: " + err.Error(),
		})
		return
	}
	
	// 获取当前用户ID
	userID := GetCurrentUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   "unauthorized",
			Message: "无法获取用户信息",
		})
		return
	}
	
	// 验证文章是否存在
	var post Post
	if err := db.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "post_not_found",
			Message: "文章不存在",
		})
		return
	}
	
	// 创建评论
	comment := Comment{
		Content: req.Content,
		UserID:  userID,
		PostID:  post.ID,
	}
	
	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "database_error",
			Message: "创建评论失败: " + err.Error(),
		})
		return
	}
	
	// 获取用户信息用于响应
	var user User
	db.First(&user, userID)
	
	// 成功响应
	c.JSON(http.StatusCreated, CommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		UserID:    comment.UserID,
		Username:  user.Username,
		PostID:    comment.PostID,
		CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}

// GetComments 获取文章评论列表处理器
func GetComments(c *gin.Context) {
	postID := c.Param("id")
	
	// 验证文章是否存在
	var post Post
	if err := db.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "post_not_found",
			Message: "文章不存在",
		})
		return
	}
	
	// 获取评论列表，预加载用户信息
	var comments []Comment
	if err := db.Where("post_id = ?", postID).Preload("User").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "database_error",
			Message: "获取评论列表失败: " + err.Error(),
		})
		return
	}
	
	// 转换为响应格式
	var commentResponses []CommentResponse
	for _, comment := range comments {
		commentResponses = append(commentResponses, CommentResponse{
			ID:        comment.ID,
			Content:   comment.Content,
			UserID:    comment.UserID,
			Username:  comment.User.Username,
			PostID:    comment.PostID,
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	
	// 成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "获取评论列表成功",
		"data":    commentResponses,
		"count":   len(commentResponses),
	})
}