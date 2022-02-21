package main

import (
	"database/sql"
	"fmt"
	//这里记得要匿名引用，因为虽然在这个库没有在main包中使用，但是在其他地方会被调用，因此必须要匿名导入
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("HelloWorld")
	// 首先编辑一个连接数据库的信息
	// 格式 用户名:密码@连接协议(address:port)/连接数据库
	dsn := "root:redhat@tcp(127.0.0.1:3306)/studygin"
	// 调用sql函数提供的open函数，指定协议为mysql，第二个参数指定连接数据库信息
	db, err := sql.Open("mysql", dsn)
	// 判断下错误，出问题就panic
	if err != nil {
		fmt.Println("连接数据库失败")
		panic(err)
	}
	// 在程序执行完毕后释放这个连接
	// 记得写在连接数据库状态判断后，防止因为无法连接到数据库导致没有db对象进而导致关闭一个空对象并报错
	defer func(db *sql.DB) {
		err := db.Close()
		// 执行个判断，如果无法关闭就panic
		if err != nil {
			fmt.Println("无法释放连接，panic")
			panic(err)
		}
	}(db)
	// 做个心跳，防止没连接上
	// 这里是因为你虽然传入了信息，并且调用了open，但是这里只是保存了连接，没有测试连接（实际连接），因此最好ping一下
	err = db.Ping()
	if err != nil {
		fmt.Println("与数据库心跳失败")
		return
	}
	fmt.Println("连接成功")

}
