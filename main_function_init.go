package main

import "fmt"
//init函数

/* 
实现包级别的一些初始化操作

1. init函数先于main函数自动执行
2. init函数没有参数和返回值
3. init函数可以有多个
4. init函数不可以被调用

执行顺序：
	初始化变量 > 初始化函数 > main函数  
*/

var i int = initVar()

func initVar() int {
	fmt.Println("initfunc...")
	return 12
}

func init() {
	fmt.Println("init...")
}


func mainIn() {
	fmt.Println("main...")	
}