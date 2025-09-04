// go开发基础作业3
package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// 在本地localhost:3306上的mysql，账号和密码都是root,有一个web3_homework数据库，
// 有一个名为students的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

func insertStudent(db *sql.DB, name string, age int, grade string) {
	res, err := db.Exec("INSERT INTO students (name, age, grade) VALUES (?, ?, ?)", name, age, grade)
	if err != nil {
		log.Fatal(err)
	}
	if id, err := res.LastInsertId(); err == nil {
		log.Printf("insert ok, id=%d", id)
	}
}

func queryStudents(db *sql.DB, age int) {
	rows, err := db.Query("SELECT id, name, age, grade FROM students WHERE age > ?", age)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id    int64
			name  string
			ageV  int
			grade string
		)
		if err := rows.Scan(&id, &name, &ageV, &grade); err != nil {
			log.Fatal(err)
		}
		log.Printf("student: id=%d name=%s age=%d grade=%s", id, name, ageV, grade)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func updateStudent(db *sql.DB, name string, grade string) {
	res, err := db.Exec("UPDATE students SET grade = ? WHERE name = ?", grade, name)
	if err != nil {
		log.Fatal(err)
	}
	if affected, err := res.RowsAffected(); err == nil {
		log.Printf("update ok, rows=%d", affected)
	}
}

func deleteStudents(db *sql.DB, age int) {
	res, err := db.Exec("DELETE FROM students WHERE age < ?", age)
	if err != nil {
		log.Fatal(err)
	}
	if affected, err := res.RowsAffected(); err == nil {
		log.Printf("delete ok, rows=%d", affected)
	}
}

// 初始化两条测试账户数据（幂等）
func seedAccounts(db *sql.DB) {
	// 账户1：余额1000.00
	if _, err := db.Exec(
		"INSERT INTO accounts (id, balance) VALUES (?, ?) ON DUPLICATE KEY UPDATE balance=VALUES(balance)",
		1, 1000.00,
	); err != nil {
		log.Fatal(err)
	}
	// 账户2：余额100.00
	if _, err := db.Exec(
		"INSERT INTO accounts (id, balance) VALUES (?, ?) ON DUPLICATE KEY UPDATE balance=VALUES(balance)",
		2, 100.00,
	); err != nil {
		log.Fatal(err)
	}
	log.Printf("seed accounts ok: (1,1000.00), (2,100.00)")
}

// 在本地localhost:3306上的mysql，账号和密码都是root,有一个web3_homework数据库，
// 有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

func transfer(db *sql.DB, fromAccountID, toAccountID int64, amount float64) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	var fromBalance float64
	if err := tx.QueryRow("SELECT balance FROM accounts WHERE id = ? FOR UPDATE", fromAccountID).Scan(&fromBalance); err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}

	if fromBalance < amount {
		_ = tx.Rollback()
		log.Fatalf("insufficient funds: have %.2f, need %.2f", fromBalance, amount)
	}

	if _, err := tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromAccountID); err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}

	if _, err := tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toAccountID); err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}

	if _, err := tx.Exec("INSERT INTO transactions (from_account_id, to_account_id, amount) VALUES (?, ?, ?)", fromAccountID, toAccountID, amount); err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	log.Printf("transfer ok: from=%d to=%d amount=%.2f", fromAccountID, toAccountID, amount)
}

// Sqlx入门
// 题目1：使用SQL扩展库进行查询
// 在本地localhost:3306上的mysql，账号和密码都是root,有一个web3_homework数据库，
// 并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 使用Sqlx连接到以上数据库，
// 要求 ：
// 先创建测试数据，插入10条数据，部门为 "技术部" 的员工信息，工资为 10000 到 20000 之间。
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
// 题目2：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/web3_homework?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// insertStudent(db, "张三", 20, "三年级")
	// queryStudents(db, 18)
	// updateStudent(db, "张三", "四年级")
	// deleteStudents(db, 15)

	// 示例：从账户1向账户2转账100.00
	seedAccounts(db)

	transfer(db, 1, 2, 100.00)

	// 运行 SQLx 作业
	runSQLxHomework()

	// 运行 GORM 作业
	runGormHomework()
}
