package main

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDatabase 初始化数据库连接
func InitDatabase() {
	dsn := "root:root@tcp(localhost:3306)/web3_homework?charset=utf8mb4&parseTime=true&loc=Local"
	
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	
	fmt.Println("数据库连接成功")
	
	// 自动迁移模式，确保数据库结构与模型同步
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}
	
	fmt.Println("数据库迁移完成")
}