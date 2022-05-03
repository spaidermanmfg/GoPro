//Go接口
//接口：定义规范，但并不去实现
/* 
定义接口语法：
type interfaceName interface {
	methodName1(parameter1 type) returnType1
}
一个接口可以被多个类型实现
*/

package main

import "fmt"


//定义接口
type USB interface {

	//定义方法
	read()
	writer()
}

//定义结构体
type Computer struct {
	name string
}

//实现接口
func (computer Computer) read() {
	fmt.Printf("%v, read...\n", computer.name)
}

//实现接口
func (computer Computer) writer() {
	fmt.Printf("%v, writer...\n", computer.name)
}


type Mobile struct {
	name string
}

func (mobile Mobile) read() {
	fmt.Printf("%v, read...\n", mobile.name)
}

func (mobile Mobile) writer() {
	fmt.Printf("%v, writer...\n", mobile.name)
}


func main3() {
	//实例化结构体
	com := Computer {
		name: "Apple",
	}
	//调用接口方法
	com.read()
	com.writer()


	mob := Mobile {
		name : "huawei",
	}
	mob.read()
	mob.writer()
}
