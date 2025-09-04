// GORM 进阶作业
package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID         uint   `gorm:"primaryKey"`
	Username   string `gorm:"uniqueIndex;size:50;not null"`
	Email      string `gorm:"uniqueIndex;size:100;not null"`
	PostsCount int    `gorm:"default:0"` // 文章数量统计字段
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	// 关联关系
	Posts []Post `gorm:"foreignKey:UserID"`
}

// Post 文章模型
type Post struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"size:200;not null"`
	Content       string `gorm:"type:text"`
	Status        string `gorm:"size:20;default:'published'"` // 文章状态
	CommentStatus string `gorm:"size:20;default:'有评论'"`       // 评论状态
	CommentsCount int    `gorm:"default:0"`                   // 评论数量统计
	UserID        uint   `gorm:"not null"`                    // 外键
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`

	// 关联关系
	User     User      `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}

// Comment 评论模型
type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	PostID    uint   `gorm:"not null"` // 外键
	UserID    uint   `gorm:"not null"` // 评论者ID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// 关联关系
	Post Post `gorm:"foreignKey:PostID"`
	User User `gorm:"foreignKey:UserID"`
}

// Post 模型的钩子函数：创建后自动更新用户的文章数量
func (p *Post) AfterCreate(tx *gorm.DB) error {
	// 更新用户的文章数量
	return tx.Model(&User{}).Where("id = ?", p.UserID).Update("posts_count", gorm.Expr("posts_count + 1")).Error
}

// Post 模型的钩子函数：删除后自动更新用户的文章数量
func (p *Post) AfterDelete(tx *gorm.DB) error {
	// 更新用户的文章数量
	return tx.Model(&User{}).Where("id = ?", p.UserID).Update("posts_count", gorm.Expr("posts_count - 1")).Error
}

// Comment 模型的钩子函数：创建后自动更新文章的评论数量
func (c *Comment) AfterCreate(tx *gorm.DB) error {
	// 更新文章的评论数量
	return tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comments_count", gorm.Expr("comments_count + 1")).Error
}

// Comment 模型的钩子函数：删除后检查并更新文章的评论状态
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 更新文章的评论数量
	if err := tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comments_count", gorm.Expr("comments_count - 1")).Error; err != nil {
		return err
	}

	// 检查评论数量，如果为0则更新评论状态
	var commentsCount int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentsCount).Error; err != nil {
		return err
	}

	if commentsCount == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论").Error
	}

	return nil
}

// 初始化数据库连接
func initDB() *gorm.DB {
	dsn := "root:root@tcp(localhost:3306)/web3_homework?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	return db
}

// 创建数据库表
func createTables(db *gorm.DB) {
	log.Println("开始创建数据库表...")

	// 自动迁移，创建表结构
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		log.Fatal("创建表失败:", err)
	}

	log.Println("数据库表创建成功")
}

// 插入测试数据
func insertTestData(db *gorm.DB) {
	log.Println("开始插入测试数据...")

	// 创建用户（使用 FirstOrCreate 避免重复）
	users := []User{
		{Username: "alice", Email: "alice@example.com"},
		{Username: "bob", Email: "bob@example.com"},
		{Username: "charlie", Email: "charlie@example.com"},
	}

	for _, user := range users {
		var existingUser User
		if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&user).Error; err != nil {
					log.Printf("创建用户失败: %v", err)
				} else {
					log.Printf("创建用户: %s", user.Username)
				}
			} else {
				log.Printf("查询用户失败: %v", err)
			}
		} else {
			log.Printf("用户已存在: %s", existingUser.Username)
		}
	}

	// 创建文章
	posts := []Post{
		{Title: "Go语言入门", Content: "Go语言是一门现代化的编程语言...", UserID: 1},
		{Title: "GORM使用指南", Content: "GORM是Go语言最流行的ORM库...", UserID: 1},
		{Title: "数据库设计原则", Content: "良好的数据库设计是系统成功的关键...", UserID: 2},
		{Title: "微服务架构", Content: "微服务架构是现代应用开发的主流...", UserID: 2},
		{Title: "Docker容器化", Content: "Docker让应用部署变得简单...", UserID: 3},
	}

	for _, post := range posts {
		if err := db.Create(&post).Error; err != nil {
			log.Printf("创建文章失败: %v", err)
		} else {
			log.Printf("创建文章: %s", post.Title)
		}
	}

	// 创建评论
	comments := []Comment{
		{Content: "写得很好，学到了很多！", PostID: 1, UserID: 2},
		{Content: "感谢分享，期待更多内容", PostID: 1, UserID: 3},
		{Content: "GORM确实很方便", PostID: 2, UserID: 3},
		{Content: "数据库设计很重要", PostID: 3, UserID: 1},
		{Content: "微服务架构有挑战性", PostID: 4, UserID: 1},
		{Content: "Docker让部署变简单了", PostID: 5, UserID: 2},
		{Content: "容器化是趋势", PostID: 5, UserID: 1},
	}

	for _, comment := range comments {
		if err := db.Create(&comment).Error; err != nil {
			log.Printf("创建评论失败: %v", err)
		} else {
			content := comment.Content
			if len(content) > 20 {
				content = content[:20] + "..."
			}
			log.Printf("创建评论: %s", content)
		}
	}

	log.Println("测试数据插入完成")
}

// 查询某个用户发布的所有文章及其对应的评论信息
func queryUserPostsWithComments(db *gorm.DB, userID uint) {
	log.Printf("查询用户 %d 的所有文章及评论:", userID)

	var user User
	if err := db.Preload("Posts.Comments").First(&user, userID).Error; err != nil {
		log.Printf("查询用户失败: %v", err)
		return
	}

	fmt.Printf("用户: %s (%s)\n", user.Username, user.Email)
	fmt.Printf("文章数量: %d\n", user.PostsCount)

	for _, post := range user.Posts {
		fmt.Printf("\n文章: %s\n", post.Title)
		content := post.Content
		if len(content) > 50 {
			content = content[:50] + "..."
		}
		fmt.Printf("  内容: %s\n", content)
		fmt.Printf("  评论数量: %d\n", post.CommentsCount)
		fmt.Printf("  评论状态: %s\n", post.CommentStatus)

		for _, comment := range post.Comments {
			fmt.Printf("    评论: %s (用户ID: %d)\n", comment.Content, comment.UserID)
		}
	}
}

// 查询评论数量最多的文章信息
func queryMostCommentedPost(db *gorm.DB) {
	log.Println("查询评论数量最多的文章:")

	var post Post
	if err := db.Preload("User").Preload("Comments").Order("comments_count DESC").First(&post).Error; err != nil {
		log.Printf("查询文章失败: %v", err)
		return
	}

	fmt.Printf("文章标题: %s\n", post.Title)
	fmt.Printf("作者: %s\n", post.User.Username)
	fmt.Printf("评论数量: %d\n", post.CommentsCount)
	fmt.Printf("评论状态: %s\n", post.CommentStatus)

	fmt.Println("评论列表:")
	for _, comment := range post.Comments {
		fmt.Printf("  - %s (用户ID: %d)\n", comment.Content, comment.UserID)
	}
}

// 测试钩子函数：删除评论
func testHooks(db *gorm.DB) {
	log.Println("测试钩子函数：删除评论")

	// 查询一篇文章的评论
	var post Post
	if err := db.Preload("Comments").First(&post).Error; err != nil {
		log.Printf("查询文章失败: %v", err)
		return
	}

	fmt.Printf("删除前 - 文章: %s, 评论数量: %d, 评论状态: %s\n",
		post.Title, post.CommentsCount, post.CommentStatus)

	// 删除所有评论
	if err := db.Where("post_id = ?", post.ID).Delete(&Comment{}).Error; err != nil {
		log.Printf("删除评论失败: %v", err)
		return
	}

	// 重新查询文章
	if err := db.First(&post, post.ID).Error; err != nil {
		log.Printf("重新查询文章失败: %v", err)
		return
	}

	fmt.Printf("删除后 - 文章: %s, 评论数量: %d, 评论状态: %s\n",
		post.Title, post.CommentsCount, post.CommentStatus)
}

func runGormHomework() {
	log.Println("=== GORM 进阶作业 ===")

	// 初始化数据库连接
	db := initDB()

	// 题目1：模型定义和表创建
	log.Println("\n--- 题目1：模型定义和表创建 ---")
	createTables(db)
	insertTestData(db)

	// 题目2：关联查询
	log.Println("\n--- 题目2：关联查询 ---")

	// 查询用户1的所有文章及评论
	queryUserPostsWithComments(db, 1)

	// 查询评论数量最多的文章
	queryMostCommentedPost(db)

	// 题目3：钩子函数测试
	log.Println("\n--- 题目3：钩子函数测试 ---")
	testHooks(db)

	log.Println("\n=== GORM 作业完成 ===")
}
