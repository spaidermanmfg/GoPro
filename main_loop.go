package main

//for循环

import (
	"fmt"
)

func m1() {
	//for结构体

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d\n", i)
	}
}

func m2() {
	//初始语句和结束语句都可以省略，初始语句可以写在循环体外部
	i := 0 //初始语句
	for i <= 10 {
		fmt.Printf("%d\n", i)
		i++ //结束语句
	}
}

func m3() {
	for {
		fmt.Println("一直执行...")
	}
}

func m4() {
	/*
		Go可以使用for range循环遍历数组，map，切片，字符串，通道
		数组、切片、字符串返回索引和数值
		map返回键和值
		通道返回值和索引
	*/

	//

	//数组
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Printf("arr[%d]=%d\n", i, v)
	}
	//map
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Printf("m[%s]=%d\n", k, v)
	}
	//切片
	var s = []int{1, 2, 3, 4, 5}
	for i, v := range s {
		fmt.Printf("s[%d]=%d\n", i, v)
	}
	//字符串
	var str = "hello world"
	for i, v := range str {
		fmt.Printf("str[%d]=%c\n", i, v)
	}
	//通道
	var ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	for v := range ch {
		fmt.Printf("ch[%d]=%d\n", v, v)
	}
}

func LoopArray() {
	//循环数组
	//数组名 = [长度]类型{数组初始化值}
	var a = [5]int{6, 2, 9, 3, 5}
	//i为索引，v为数组元素,不使用索引可以使用_代替
	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}
}

func LoopSlice() {
	//循环切片
	//切片名 = [长度]类型，切片是一个动态数组，可以改变长度，但是不能改变容量
	var s = []int{3, 5, 4, 6, 7}
	for _, v := range s {
		fmt.Printf("%d\n", v)
	}
}

func LoopMap() {
	//循环map
	//map名 = map[键类型]值类型{键值对}
	var m = map[string]string{}
	m["name"] = "Lily"
	m["age"] = "18"
	m["email"] = "afasda@gemail.com"
	m["adress"] = "Beijing"
	for key, value := range m {
		fmt.Printf("%s : %s\n", key, value)
	}
}

func LoopString() {
	//循环字符串
	var str string = "hello world"
	for i, v := range str {
		fmt.Printf("str[%d] = %c\n", i, v)
	}
}

func maint() {
	/*
		for 初始语句; 条件表达式; 循环语句 {
			//循环体语句
		}
	*/

	//m1()
	//m2()
	//m3()
	LoopArray()
	LoopSlice()
	LoopMap()
	LoopString()
}
