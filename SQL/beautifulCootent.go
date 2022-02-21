package main

import (
	"database/sql"
	"fmt"
	//这里记得要匿名引用，因为虽然在这个库没有在main包中使用，但是在其他地方会被调用，因此必须要匿名导入
	_ "github.com/go-sql-driver/mysql"
)

// 先全局声明一个数据库对象,用指针
var db *sql.DB

// 创建一个初始化函数，用于初始化这个数据库对象
func init() {
	// 在初始化函数中创建一个函数全局错误信息
	var err error
	// 数据库连接信息
	dsn := "root:redhat@tcp(127.0.0.1:3306)/studygin"
	// 创建对象连接
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("创建数据库连接失败，错误信息请查看Panic")
		panic(err)
	}
	fmt.Println("创建连接对象成功，正在尝试ping")
	// 做一个ping处理，看下是否真的连接上了
	err = db.Ping()
	if err != nil {
		fmt.Println("虽然与数据库的连接已经被创建，但是并未被正确的连接，具体信息请查看错误代码")
		panic(err)
	}
	// 设置下最大连接数，这个取决于个人需求，如果连接数超过了会被闲置
	db.SetConnMaxLifetime(200)
	// 设置下最大空闲连接数，取决于个人，就是最低空闲连接数，无论什么状态都至少保存这些连接数
	db.SetMaxIdleConns(30)

}
func main() {
	fmt.Println("Hello")
	err := db.Ping()
	if err != nil {
		fmt.Println("虽然与数据库的连接已经被创建，但是并未被正确的连接，具体信息请查看错误代码")
		panic(err)
	}
	fmt.Println("连接成功")
}
