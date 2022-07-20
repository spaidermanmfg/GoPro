//类型定义
package main

import "fmt"

/*
类型定义语法： type varName varType
类型别名语法： type varName = varType
类型定义和类型别名的区别：
	1.类型定义定义了一个全新的类型，与之前的类型不同；类型别名只是使用一个别名来替代之前的类型。
	2.类型别名只会在代码中存在，编译后不会存在
	3.类型别名可以调用原来类型的方法和属性，但是重新定义的类型定义不可以调用之前类型的任何方法。
*/

func typeDefine() {
	//类型定义
	type myInt int
	var a myInt = 1
	fmt.Printf("%T , %v\n", a, a) //main.myInt , 1
}

func typeOtherName() {
	//类型别名
	type myInt = int
	var a myInt = 1
	fmt.Printf("%T , %v\n", a, a) //int , 1

}

func mainType() {
	typeDefine()
	typeOtherName()
}
