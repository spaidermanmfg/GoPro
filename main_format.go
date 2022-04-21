package main
//Go格式化输出
import (
	"fmt"
)

type Person struct {
	Name string
}
func mainm() {
	/*
	Printf: Print Format
	*/
	name := Person{Name: "Mark"}
	//%v var 变量
	fmt.Printf("%v\n", name)
	//%#v var 变量,带包名方法名
	fmt.Printf("%#v\n", name)
	//%T Type 类型
	fmt.Printf("%T\n", name)
	//%% 无特殊含义 %
	fmt.Printf("%%\n")
	//%t bool 布尔值
	fmt.Printf("%t\n", true)
	fmt.Printf("--------------------------\n")
	i := 8
	//%d int 十进制
	fmt.Printf("%d\n", i)
	//%b binary 二进制
	fmt.Printf("%b\n", i)
	//%o octal 八进制 
	fmt.Printf("%o\n", i)
	//%x hex 十六进制
	fmt.Printf("%x\n", i)//小写
	fmt.Printf("%X\n", i)//大写
	//符号格式
	i = 91
	fmt.Printf("%c\n", i)//[
	//%s string
	fmt.Printf("%s\n", name.Name)
	//%p pointer 指针
	fmt.Printf("%p\n", &name)
	fmt.Printf("--------------------------\n")
}