package main

import "fmt"

//function type

func sum1(a int, b int) int {
	return a + b
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func mainee() {

	//使用type关键字定义函数类型，函数类型可以作为参数传递，用来模拟类型相同的函数
	//type type_name func (parameter_list) return_value
	type fun func(int, int) int

	var f fun
	f = sum1
	sum := f(6, 7)
	fmt.Printf("sum = %v\n", sum)

	f = max
	max := f(6, 7)
	fmt.Printf("max = %v\n", max)
}
