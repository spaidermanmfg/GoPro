package main

import (
	"fmt"
)

func main() {
	index := 6
	double(&index)
	fmt.Println(index)//获取变量地址
	
	myInt := 4
	myIntPoint := &myInt
	*myIntPoint = 8//更改指针处的值
	fmt.Println(myIntPoint)//变量地址
	fmt.Println(*myIntPoint)//指针引用的值

	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate
	(&lies)
	fmt.Println(lies)

}

func double (num *int) {
	*num *= 2
}

func negate (myBoolean *bool){
	*myBoolean = !*myBoolean
}