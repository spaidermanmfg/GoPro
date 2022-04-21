package main

import (
	"fmt"
)


//指针
/* 
每个变量在运行时在内存中都有一个地址，这个地址代表变量在内存中的位置。
不光可以通过变量名来访问变量，还可以通过变量的地址来访问变量。
& 取地址符号，取变量的地址。
* 根据地址取值。

声明语法
var varname *vartype

*/

func intPoint() {
	//声明int类型指针
	var ip *int
	fmt.Printf("%v\n", ip)
	fmt.Printf("%v\n", &ip)

	//声明变量
	var i int = 100
	//为指针赋值
	ip = &i
	//输出
	fmt.Printf("ip: %v\n", &ip)//取地址
	fmt.Printf("ip: %v\n", *ip)//取值
}

func stringPoint() {
	//声明string类型指针
	var sp *string
	var str string = "Model"
	sp = &str
	fmt.Printf("spAddress: %v | spValue: %v\n", &sp, *sp)
}

func boolPoint() {
	//声明bool类型指针
	var bp *bool
	var b bool = true
	bp = &b
	fmt.Printf("bpAddress: %v | bpValue: %v\n", &bp, *bp)
}

func pointArray() {
	//指针数组，数组中存储指针类型
	var a = [5]int{1,2,3,4,5}
	var pa [5]*int
	fmt.Printf("pa: %v\n", pa)//[<nil> <nil> <nil> <nil> <nil>]

	//整数地址赋值给指针数组
	for i := 0; i < len(a); i++ {
		pa[i] = &a[i]
	}

	fmt.Printf("pa: %v\n", pa)//[0xc0000aa030 0xc0000aa038 0xc0000aa040 0xc0000aa048 0xc0000aa050]

	for i := 0; i < len(pa); i++ {
		fmt.Printf("pa[%v]: %v\n", i, *pa[i])//pa[0]: 1 pa[1]: 2 pa[2]: 3 pa[3]: 4 pa[4]: 5
	}
}

func mainPoint() {
	intPoint()
	stringPoint()
	boolPoint()	
	pointArray()
}