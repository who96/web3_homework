// SQLx入门作业
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Employee 结构体对应 employees 表
type Employee struct {
	ID         int64   `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// Book 结构体对应 books 表
type Book struct {
	ID     int64   `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

// 创建 books 表
func createBooksTable(db *sqlx.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS books (
		id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(200) NOT NULL,
		author VARCHAR(100) NOT NULL,
		price DECIMAL(10,2) NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`

	if _, err := db.Exec(query); err != nil {
		log.Fatal("创建 books 表失败:", err)
	}
	log.Println("books 表创建成功")
}

// 插入10条技术部员工测试数据
func insertTechEmployees(db *sqlx.DB) {
	names := []string{"张三", "李四", "王五", "赵六", "钱七", "孙八", "周九", "吴十", "郑十一", "王十二"}

	for i, name := range names {
		// 生成10000到20000之间的随机工资
		salary := 10000.0 + rand.Float64()*10000.0

		query := `INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)`
		if _, err := db.Exec(query, name, "技术部", salary); err != nil {
			log.Fatal("插入员工数据失败:", err)
		}
		log.Printf("插入员工 %d: %s, 技术部, %.2f", i+1, name, salary)
	}
}

// 查询技术部所有员工
func queryTechEmployees(db *sqlx.DB) []Employee {
	var employees []Employee
	query := `SELECT id, name, department, salary FROM employees WHERE department = ?`

	if err := db.Select(&employees, query, "技术部"); err != nil {
		log.Fatal("查询技术部员工失败:", err)
	}

	log.Printf("查询到 %d 名技术部员工:", len(employees))
	for _, emp := range employees {
		log.Printf("  ID: %d, 姓名: %s, 部门: %s, 薪资: %.2f",
			emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	return employees
}

// 查询工资最高的员工
func queryHighestPaidEmployee(db *sqlx.DB) Employee {
	var employee Employee
	query := `SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1`

	if err := db.Get(&employee, query); err != nil {
		log.Fatal("查询最高薪资员工失败:", err)
	}

	log.Printf("工资最高的员工: ID: %d, 姓名: %s, 部门: %s, 薪资: %.2f",
		employee.ID, employee.Name, employee.Department, employee.Salary)

	return employee
}

// 插入一些测试书籍数据
func insertTestBooks(db *sqlx.DB) {
	books := []Book{
		{Title: "Go语言圣经", Author: "Alan Donovan", Price: 89.00},
		{Title: "深入理解计算机系统", Author: "Randal E. Bryant", Price: 128.00},
		{Title: "算法导论", Author: "Thomas H. Cormen", Price: 158.00},
		{Title: "设计模式", Author: "Gang of Four", Price: 45.00},
		{Title: "重构", Author: "Martin Fowler", Price: 78.00},
		{Title: "代码整洁之道", Author: "Robert C. Martin", Price: 65.00},
		{Title: "人月神话", Author: "Frederick P. Brooks", Price: 35.00},
		{Title: "程序员修炼之道", Author: "Andrew Hunt", Price: 55.00},
	}

	for _, book := range books {
		query := `INSERT INTO books (title, author, price) VALUES (?, ?, ?)`
		if _, err := db.Exec(query, book.Title, book.Author, book.Price); err != nil {
			log.Fatal("插入书籍数据失败:", err)
		}
		log.Printf("插入书籍: %s, %s, %.2f", book.Title, book.Author, book.Price)
	}
}

// 查询价格大于50元的书籍
func queryExpensiveBooks(db *sqlx.DB) []Book {
	var books []Book
	query := `SELECT id, title, author, price FROM books WHERE price > ? ORDER BY price DESC`

	if err := db.Select(&books, query, 50.0); err != nil {
		log.Fatal("查询高价书籍失败:", err)
	}

	log.Printf("查询到 %d 本价格大于50元的书籍:", len(books))
	for _, book := range books {
		log.Printf("  ID: %d, 书名: %s, 作者: %s, 价格: %.2f",
			book.ID, book.Title, book.Author, book.Price)
	}

	return books
}

func runSQLxHomework() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 连接数据库
	db, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/web3_homework?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	defer db.Close()

	log.Println("=== SQLx 入门作业 ===")

	// 题目1：Employee 相关操作
	log.Println("\n--- 题目1：Employee 操作 ---")

	// 插入10条技术部员工数据
	log.Println("1. 插入10条技术部员工数据:")
	insertTechEmployees(db)

	// 查询技术部所有员工
	log.Println("\n2. 查询技术部所有员工:")
	techEmployees := queryTechEmployees(db)
	fmt.Printf("技术部员工总数: %d\n", len(techEmployees))

	// 查询工资最高的员工
	log.Println("\n3. 查询工资最高的员工:")
	highestPaid := queryHighestPaidEmployee(db)
	fmt.Printf("最高薪资员工: %s, %.2f元\n", highestPaid.Name, highestPaid.Salary)

	// 题目2：Book 相关操作
	log.Println("\n--- 题目2：Book 操作 ---")

	// 创建 books 表
	log.Println("1. 创建 books 表:")
	createBooksTable(db)

	// 插入测试书籍数据
	log.Println("\n2. 插入测试书籍数据:")
	insertTestBooks(db)

	// 查询价格大于50元的书籍
	log.Println("\n3. 查询价格大于50元的书籍:")
	expensiveBooks := queryExpensiveBooks(db)
	fmt.Printf("高价书籍总数: %d\n", len(expensiveBooks))

	log.Println("\n=== 作业完成 ===")
}
