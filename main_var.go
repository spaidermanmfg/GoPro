package main

//调用包
import (
	"fmt"
	"GoPro/user"
)
 
//Go声明变量
func mains() {
	s := user.Hello()
	fmt.Printf("s: %v", s)
	//声明变量
	var name string
	var age int
	//类型推断
	var height = 1.75
	//变量初始化
	var polo bool = true
	//赋值
	name = "张三"
	age = 18
	//输出
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("height: %v\n", height)
	fmt.Printf("polo: %v\n", polo)
	
}