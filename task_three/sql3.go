package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // 使用 SQLite 作为示例
)

// Employee 结构体对应 employees 表
type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

// Book 结构体对应 books 表
type Book struct {
	ID     int    `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Price  int    `db:"price"`
}

// 题目1: 查询技术部所有员工
func queryTechnicalEmployees(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee
	query := `SELECT id, name, department, salary FROM employees WHERE department = ?`

	err := db.Select(&employees, query, "技术部")
	if err != nil {
		return nil, fmt.Errorf("查询技术部员工失败: %v", err)
	}

	return employees, nil
}

// 题目1: 查询工资最高的员工
func queryHighestSalaryEmployee(db *sqlx.DB) (Employee, error) {
	var employee Employee
	query := `SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1`

	err := db.Get(&employee, query)
	if err != nil {
		return Employee{}, fmt.Errorf("查询最高工资员工失败: %v", err)
	}

	return employee, nil
}

// 题目2: 查询价格大于50元的书籍
func queryExpensiveBooks(db *sqlx.DB) ([]Book, error) {
	var books []Book
	query := `SELECT id, title, author, price FROM books WHERE price > ? ORDER BY price DESC`

	err := db.Select(&books, query, 50)
	if err != nil {
		return nil, fmt.Errorf("查询高价书籍失败: %v", err)
	}

	return books, nil
}

// 创建示例数据
func initDB(db *sqlx.DB) error {
	// 创建 employees 表
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS employees (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			department TEXT NOT NULL,
			salary INTEGER NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	// 创建 books 表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			author TEXT NOT NULL,
			price INTEGER NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	// 插入示例员工数据
	employees := []Employee{
		{Name: "张三", Department: "技术部", Salary: 8000},
		{Name: "李四", Department: "技术部", Salary: 9000},
		{Name: "王五", Department: "市场部", Salary: 7000},
		{Name: "赵六", Department: "技术部", Salary: 8500},
	}

	for _, emp := range employees {
		_, err := db.Exec(`INSERT INTO employees (name, department, salary) VALUES (?, ?, ?)`,
			emp.Name, emp.Department, emp.Salary)
		if err != nil {
			return err
		}
	}

	// 插入示例书籍数据
	books := []Book{
		{Title: "Go语言编程", Author: "作者A", Price: 65},
		{Title: "数据库设计", Author: "作者B", Price: 45},
		{Title: "算法导论", Author: "作者C", Price: 89},
		{Title: "Web开发", Author: "作者D", Price: 55},
		{Title: "Python入门", Author: "作者E", Price: 35},
	}

	for _, book := range books {
		_, err := db.Exec(`INSERT INTO books (title, author, price) VALUES (?, ?, ?)`,
			book.Title, book.Author, book.Price)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// 连接到SQLite数据库
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	defer db.Close()

	// 初始化数据库和示例数据
	if err := initDB(db); err != nil {
		log.Fatal("初始化数据库失败:", err)
	}

	fmt.Println("=== 题目1：员工查询 ===")

	// 查询技术部所有员工
	techEmployees, err := queryTechnicalEmployees(db)
	if err != nil {
		log.Printf("错误: %v", err)
	} else {
		fmt.Println("技术部所有员工:")
		for _, emp := range techEmployees {
			fmt.Printf("  ID: %d, 姓名: %s, 部门: %s, 工资: %d\n",
				emp.ID, emp.Name, emp.Department, emp.Salary)
		}
	}

	// 查询工资最高的员工
	highestSalaryEmp, err := queryHighestSalaryEmployee(db)
	if err != nil {
		log.Printf("错误: %v", err)
	} else {
		fmt.Printf("\n工资最高的员工: ID: %d, 姓名: %s, 部门: %s, 工资: %d\n",
			highestSalaryEmp.ID, highestSalaryEmp.Name, highestSalaryEmp.Department, highestSalaryEmp.Salary)
	}

	fmt.Println("\n=== 题目2：书籍查询 ===")

	// 查询价格大于50元的书籍
	expensiveBooks, err := queryExpensiveBooks(db)
	if err != nil {
		log.Printf("错误: %v", err)
	} else {
		fmt.Println("价格大于50元的书籍:")
		for _, book := range expensiveBooks {
			fmt.Printf("  ID: %d, 书名: %s, 作者: %s, 价格: %d元\n",
				book.ID, book.Title, book.Author, book.Price)
		}
	}
}

