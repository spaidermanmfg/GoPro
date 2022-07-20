package main

import "fmt"

/*
defer关键字,类似于栈数据结构，先进后出，先定义的后执行

用于注册延迟调用，会将之后跟随的语句进行延迟处理，被defer的语句会在return前执行
用途：
	关闭文件句柄
	锁资源释放
	数据库连接释放
	资源清理
举例：
	ctrl + z
*/
func mainll() {
	fmt.Println(111)
	fmt.Println(222)
	fmt.Println(333)
	fmt.Println(444)
	fmt.Println(555)

	fmt.Println()

	fmt.Println(111)
	defer fmt.Println(222)
	defer fmt.Println(333)
	defer fmt.Println(444)
	fmt.Println(555)
}
