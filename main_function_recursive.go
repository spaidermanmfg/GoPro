package main

import "fmt"


/* 
递归的英文是 recursive，意思是反复，递归的意思是调用自身。
递归的定义是：在函数内部调用自身本身，并且返回。
递归的实现是：在函数内部调用自身本身，并且返回。
递归的应用场景是：在程序中，如果一个函数调用了自身，那么这个函数就是递归函数。


递归：函数调用自身
三要素：
    函数自己调用自己
	要有结束条件，没有就会成为死循环
	容易出现栈空间溢出
*/


//递归实现阶乘
func mult(a int) int {
	if a == 1 {
		return 1
	}else {
		return a * mult(a - 1)
	}
}

//递归实现斐波那契数列
//f(n) = f(n-1) + f(n-2), f(2) = f(1) = 1
func fibonacci(n int) int {
	if n ==1 || n ==2 {
		return 1
	}else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}


func mainre() {
	mult := mult(4)
	fmt.Printf("mult: %v\n", mult)

	fib := fibonacci(10)
	fmt.Printf("fib: %v\n", fib)
}