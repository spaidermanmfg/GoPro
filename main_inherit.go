package main

import "fmt"

//Go 继承
/*
通过嵌套结构体来实现
*/

type Animal struct {
	name string
	age  int
}

func (ani Animal) eat() {
	fmt.Println(ani.name, "恰饭。。。")
}

func (ani Animal) sleep() {
	fmt.Println(ani.name, "睡大觉。。。")
}

type Sheep struct {
	ani   Animal //可以看作继承
	color string
}

type Dargon struct {
	ani   Animal
	color string
}

func main8() {
	sheep := Sheep{
		ani: Animal{
			name: "咩咩",
			age:  3,
		},
		color: "white",
	}

	dargon := Dargon{
		ani: Animal{
			name: "long",
			age:  55,
		},
		color: "blue",
	}

	sheep.ani.eat()
	sheep.ani.sleep()
	fmt.Printf("color: %v\n", sheep.color)
	fmt.Printf("age: %v\n", sheep.ani.age)
	fmt.Printf("name: %v\n", sheep.ani.name)
	fmt.Println("-----------------")
	dargon.ani.eat()
	dargon.ani.sleep()
	fmt.Printf("color: %v\n", dargon.color)
	fmt.Printf("age: %v\n", dargon.ani.age)
	fmt.Printf("name: %v\n", dargon.ani.name)

}
