package main

import "fmt"

/*
function
function contains a function name, parameters, return value and body
1. 函数不允许重载，不能重名
2. 函数不能嵌套
3. 函数可以作为返回值返回
4. 函数可以作为参数传递
5. 函数参数可以无名
函数定义语法
func function_name(parameter_list) return_value {
	//function_body
}
*/

//行参
func sum(a int, b int) (c int) {
	//The summation of function
	c = a + b
	return c
}

func comp(a int, b int) (max int) {
	//The comparison of function
	if a > b {
		max = a
	} else {
		max = b
	}
	return max
}

func return_func() (name string, age int, hack bool) {
	//多个返回值使用逗号分隔
	//可以自己定义返回值名称覆盖原有名称
	name = "kimi"
	age = 18
	hack = false
	return //return name, age, hack等价于return
}

//parameter
//copy
func param1(a int) {
	a = 100
}

//切片、数组、map等类型的参数会被修改值
func param2(a []int) {
	a[0] = 999
}

//函数作为参数传递
func sum2(a int, b int) {
	c := a + b
	fmt.Printf("\na + b = %v\n", c)
}
func trans_func(a int, b int, f func(int, int)) {
	f(a, b)
}

//函数作为返回值
func sums(a int, b int) int {
	return a + b
}

func subs(a int, b int) int {
	return a - b
}

func back_func(s string) func(int, int) int {
	switch s {
	case "+":
		return sums
	case "-":
		return subs
	default:
		return nil
	}
}

//变长参数
func change_param(name string, yep bool, num int, args ...int) {
	fmt.Printf("name: %v, yep: %v, num: %v, args: %v\n", name, yep, num, args)
	for _, v := range args {
		fmt.Printf("%v\t", v)
	}
}

func mainrr() {
	//call function
	//实参
	result := sum(4, 5)
	fmt.Printf("result: %v\n", result)
	result = comp(4, 5)
	fmt.Printf("max: %v\n", result)
	name, age, hack := return_func()
	fmt.Printf("name: %v, age: %v, hack: %v\n", name, age, hack)

	//值没有被修改
	a := 200
	param1(a)
	fmt.Printf("a: %v\n", a)

	//切片、数组、map等类型的参数会被修改值
	b := []int{1, 2, 3, 4}
	param2(b) //值被修改
	fmt.Printf("b: %v\n", b)

	//变长参数
	change_param("mike", true, 88, 1, 2, 3, 4, 5, 6, 7, 8, 9, 99)

	//函数作为参数传递
	trans_func(6, 7, sum2)

	//函数作为返回值使用
	fc := back_func("+")
	re := fc(5, 6)
	fmt.Printf("%v\n", re)

	fc = back_func("-")
	re = fc(6, 7)
	fmt.Printf("%v\n", re)

}
