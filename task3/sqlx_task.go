package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 不要忘了导入数据库驱动
	"github.com/jmoiron/sqlx"
)

type Employees struct {
	id         int
	name       string
	department string
	salary     int
}

var mysqlDb *sqlx.DB

func initDB() (err error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=true&loc=Local"
	mysqlDb, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect db failed:", err)
		return err
	}
	mysqlDb.SetMaxIdleConns(10)
	mysqlDb.SetMaxOpenConns(20)
	return nil
}
func selectEmployees() {

	if err := initDB(); err != nil {
		fmt.Println("init db failed:", err)
		return
	}
	selectByDepartment := "SELECT id, name, department, salary FROM employees WHERE department = ?"
	var employees []Employees
	if err := mysqlDb.Get(&employees, selectByDepartment, "技术部"); err != nil {
		fmt.Println("get employees failed:", err)
	}
	fmt.Println("employees:", employees)
}
func selectEmployeesSalary() {

	if err := initDB(); err != nil {
		fmt.Println("init db failed:", err)
		return
	}
	selectByDepartment := "SELECT id, name, department, salary FROM employees ORDER BY salary desc LIMIT 1"
	var employees []Employees
	if err := mysqlDb.Get(&employees, selectByDepartment); err != nil {
		fmt.Println("get employees failed:", err)
	}
	fmt.Println("employees:", employees)
}

type Book struct {
	id     int
	title  string
	author string
	price  float64
}

func selectBook() {
	if err := initDB(); err != nil {
		fmt.Println("init db failed:", err)
		return
	}
	selectBookSql := "SELECT * FROM books WHERE price > 50"
	var book Book
	if err := mysqlDb.Get(&book, selectBookSql); err != nil {
		fmt.Println("get book failed:", err)
		return
	}
	fmt.Println("book:", book)

}
