package main

import "fmt"

//结构体嵌套，实现继承功能

type Dog struct {
	name string
	age int
	color string
}

type Persons struct {
	name string 
	age int
	dog Dog
}

func main13() {
	//初始化dog
	var dog Dog
	dog = Dog{
		name: "mark",
		age: 45,
		color: "red",
	}
	fmt.Printf("dog: %v\n", dog)

	//初始化人
	var person Persons
	person = Persons{
		name: "tom",
		age: 45,
		dog: dog,
	}

	re := person.dog.name//获取person中dog的name
	fmt.Printf("dogname: %v\n", re)
	fmt.Printf("person: %v\n", person)
}