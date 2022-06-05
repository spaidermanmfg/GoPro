package main

/*
操作数据库
1.下载安装数据库
2.导入数据库驱动
     go get -u github.com/go-sql-driver/mysql
3.导包
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
4.创建数据库连接
	db, err := sql.Open("mysql", "root:12345678@/dbname")
5.操作数据库
	//插入数据
	//更新数据
	//查询数据
	//删除数据
*/

import (
	"fmt"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id int
	username string
	password string
}


//定义全局变量db
var db *sql.DB


//插入数据
func insertData(username string, password string) {
	sqlStr := "insert into dbname (name, age) values (?, ?)"
	r , err := db.Exec(sqlStr, username, password)
	if err != nil {
		panic(err)
	}else{
		fmt.Println("插入成功")
	}
}

//查询数据
//单行查询
func queryRowDate() {
	sqlStr := "select * from dbname where id = ?"
	var user User
	err := db.QueryRow(sqlStr, 1).Scan(&user.id, &user.username, &user.password)
	if err != nil {
		fmt.Println("查询失败")
	}else{
		fmt.Println("查询成功")
	}
}

//多条数据
func queryManyDate() {
	sqlStr := "select * from dbname"
	r, err := db.Query(sqlStr)
	var user User
	defer r.Close()
	if err != nil {
		fmt.Println("查询失败")
	}else{
		for r.Next() {
			r.Scan(&user.id, &user.username, &user.password)
			fmt.Println(user)
		}
	}

}

//更新数据
func updateData() {
	sqlStr := "update dbname set username = ? where id = ?"
	res, err := db.Exec(sqlStr, "小明", 1)
	if err != nil {
		fmt.Println("更新失败")
	}else {
		fmt.Println("更新成功")
		i , _ := res.RowsAffected()
		fmt.Println(i)//更新的行数
	}
}


//删除数据
func deleteData() {
	sqlStr := "delete from dbname where id = ?"
	res, err := db.Exec(sqlStr, 1)
	if err != nil {
		fmt.Println("删除失败")
	}else {
		fmt.Println("删除成功")
		i , _ := res.RowsAffected()
		fmt.Println(i)//删除的行数
	}
}

//初始化数据库
func initDB() (err error) {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	//连接数据库
	err = db.Ping()
	if err != nil {
		fmt.Println("连接数据库失败")
		return err
	}
	return nil
}


func main() {
	// //open函数并不创建数据库连接，只是验证参数格式是否正确，使用Ping函数创建数据库连接
	// db, err := sql.Open("mysql", "roor:12345678@/dbname")
	// if err != nil {
	// 	panic(err)
	// }
	// // See "Important settings" section.
	// db.SetConnMaxLifetime(time.Minute * 3)//设置连接最大连接时间
	// db.SetMaxOpenConns(10)//最大连接数
	// db.SetMaxIdleConns(10)//最大空闲连接数

	err := initDB()
	if err != nil {
		fmt.Println("初始化数据库失败")
	}else{
		fmt.Println("初始化数据库成功")
	}
	
}