//结构体
package main

import "fmt"

/* 
定义结构体就像自定义了一个类型
结构体语法：
type structName struct {
	fieldName fieldType
	fieldName fieldType
	...
}
通过.来访问结构体成员：structName.fieldName
*/

//定义一个结构体Person
type Personal struct {
	name string
	age int
	email string
	phone int
}

type nyjah struct {
	name string
	age int
	job string
	combine bool
}

func mainst() {
	var mark Personal
	mark.name = "Mark"
	mark.age = 18
	mark.email = "mark@google.com"
	mark.phone = 110
	fmt.Printf("mark: %v\n", mark)


	//匿名结构体，可以声明在函数内部，适用于成员变量少，简单的情况
	var john struct {
		name string 
		age int
		email string
	}
	john.name = "John"
	john.age = 13
	john.email = "john@google.com"
	fmt.Printf("john: %v\n", john)


	//初始化结构体
	var hoston nyjah
	//1.键值对初始化
	hoston = nyjah {
		name: "nyjah",
		age: 11,
		job: "skate",
		combine: false,
	}
	fmt.Printf("hoston: %v\n", hoston)
	//2.顺序初始化
	var lador nyjah
	lador = nyjah {
		"ladon",
		25,
		"f1 driver",
		true,
	}
	fmt.Printf("lador: %v\n", lador)
	//3.部分初始化
	var lec nyjah 
	lec = nyjah {
		name: "lec",
		age: 23,
	}
	fmt.Printf("lec: %v\n", lec)		
}