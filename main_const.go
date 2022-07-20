package main

import "fmt"

//Go声明常量

func mainw() {

	//定义常量
	const PI float32 = 3.14
	fmt.Printf("PI: %v\n", PI)
	//使用iota关键字
	const (
		a = iota //0
		b = iota //1
		_        //2,使用下划线跳过
		c = 100  //3，中间插入值跳过
		d = iota //4
	)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
}
