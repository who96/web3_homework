package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库连接
	InitDatabase()
	
	// 创建Gin路由器
	r := gin.Default()
	
	// 基础健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "博客后端系统运行正常",
		})
	})
	
	// 用户认证相关路由
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", Register)
		auth.POST("/login", Login)
	}
	
	// 公开访问的路由
	api := r.Group("/api")
	{
		// 文章相关
		api.GET("/posts", GetPosts)
		api.GET("/posts/:id", GetPostByID)
		
		// 评论相关
		api.GET("/posts/:id/comments", GetComments)
	}
	
	// 需要认证的路由
	protected := r.Group("/api/protected", AuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			userID := GetCurrentUserID(c)
			username := GetCurrentUsername(c)
			c.JSON(http.StatusOK, gin.H{
				"message":  "访问成功",
				"user_id":  userID,
				"username": username,
			})
		})
		
		// 文章管理
		protected.POST("/posts", CreatePost)
		protected.PUT("/posts/:id", UpdatePost)
		protected.DELETE("/posts/:id", DeletePost)
		
		// 评论管理
		protected.POST("/posts/:id/comments", CreateComment)
	}
	
	// 启动服务器
	r.Run(":8080")
}