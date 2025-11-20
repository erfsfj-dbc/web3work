package main

import (
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 题目1：基本CRUD操作
// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

type Student struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"type:varchar(100)"`
	Age    int    `gorm:"type:int"`
	Grade  string `gorm:"type:varchar(100)"`
	Remark string `gorm:"type:varchar(100)"`
}

func ques01(db *gorm.DB) {
	// db, err := gorm.Open(mysql.Open("root:1010@tcp(127.0.0.1:3306)/db_go_test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	db.AutoMigrate(&Student{})
	db.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})
	s := &[]Student{}
	db.Where("age > ?", 18).Find(s)
	fmt.Println(s)
	db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
	db.Model(&Student{}).Where("id = ?", "3").Update("age", "12")
	db.Where("age < ?", 10).Delete(&Student{})
}

// 题目2：事务语句
// 假设有两个表：
// accounts 表（包含字段 id 主键， balance 账户余额）和
// transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
// 在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
// 并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
type Account struct {
	ID      uint  `gorm:"primaryKey"`
	Balance int64 `gorm:"type:bigint"`
}

type TransferRecord struct {
	ID            uint  `gorm:"primaryKey"`
	FromAccountID uint  `gorm:"index"`
	ToAccountID   uint  `gorm:"index"`
	Amount        int64 `gorm:"type:bigint"`
}

func transactionDemo(db *gorm.DB, fromAccountID, toAccountID uint, amount int64) {

	err := db.Transaction(func(tx *gorm.DB) error {
		var from Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&from, fromAccountID).Error; err != nil {
			return err
		}

		if from.Balance < amount {
			return errors.New("余额不足")
		}

		var to Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&to, toAccountID).Error; err != nil {
			return err
		}

		if err := tx.Model(&Account{}).
			Where("id = ?", fromAccountID).
			Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return err
		}

		if err := tx.Model(&Account{}).
			Where("id = ?", toAccountID).
			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}

		record := TransferRecord{
			FromAccountID: fromAccountID,
			ToAccountID:   toAccountID,
			Amount:        amount,
		}
		if err := tx.Create(&record).Error; err != nil {
			return err
		}

		// 所有操作都成功了，事务提交
		return nil
	})

	if err != nil {
		fmt.Println("事务执行失败，已回滚：", err)
	} else {
		fmt.Println("事务执行成功，已提交！")
	}
}

// Sqlx入门
// 题目1：使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
type Employee struct {
	ID         int64   `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

type Book struct {
	ID     int64   `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

// queryEmployeesByDepartment 使用 sqlx 查询指定部门的员工信息
func queryEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	var employees []Employee

	const query = `
		SELECT id, name, department, salary
		FROM employees
		WHERE department = ?
	`

	if err := db.Select(&employees, query, department); err != nil {
		return nil, err
	}

	return employees, nil
}

// queryHighestPaidEmployee 查询工资最高的员工
func queryHighestPaidEmployee(db *sqlx.DB) (Employee, error) {
	var employee Employee

	const query = `
		SELECT id, name, department, salary
		FROM employees
		ORDER BY salary DESC
		LIMIT 1
	`

	if err := db.Get(&employee, query); err != nil {
		return Employee{}, err
	}

	return employee, nil
}

// queryBooksAbovePrice 查询价格高于指定阈值的书籍，返回类型安全的映射
func queryBooksAbovePrice(db *sqlx.DB, minPrice float64) ([]Book, error) {
	var books []Book

	const query = `
		SELECT id, title, author, price
		FROM books
		WHERE price > ?
	`

	if err := db.Select(&books, query, minPrice); err != nil {
		return nil, err
	}

	return books, nil
}

func main() {
	// 先尝试加载本地配置文件，若不存在则忽略
	if err := godotenv.Load(".config.env"); err != nil {
		fmt.Println("未找到 .config.env 文件，尝试直接读取环境变量")
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		panic("环境变量 DB_DSN 未配置，请在本地配置数据库连接字符串")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&Student{}, &Account{}, &TransferRecord{}, &Employee{}, &Book{}); err != nil {
		fmt.Println("AutoMigrate error:", err)
		panic(err)
	}

	// 初始化示例账户
	var count int64
	db.Model(&Account{}).Count(&count)
	if count < 2 {
		db.Where("id IN ?", []int{1, 2}).Delete(&Account{})
		db.Create(&[]Account{
			{ID: 1, Balance: 500},
			{ID: 2, Balance: 300},
		})
	}

	transactionDemo(db, 1, 2, 250)

	sqlxDB, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("sqlx 连接失败: %v", err))
	}
	defer sqlxDB.Close()

	deptEmployees, err := queryEmployeesByDepartment(sqlxDB, "技术部")
	if err != nil {
		fmt.Println("查询技术部员工失败:", err)
	} else {
		fmt.Println("技术部员工列表:", deptEmployees)
	}

	highest, err := queryHighestPaidEmployee(sqlxDB)
	if err != nil {
		fmt.Println("查询薪资最高员工失败:", err)
	} else {
		fmt.Println("薪资最高的员工:", highest)
	}

	expensiveBooks, err := queryBooksAbovePrice(sqlxDB, 50)
	if err != nil {
		fmt.Println("查询高价图书失败:", err)
	} else {
		fmt.Println("价格大于 50 元的书籍:", expensiveBooks)
	}
}
