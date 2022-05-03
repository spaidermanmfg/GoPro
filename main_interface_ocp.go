package main

import "fmt"
//Go接口实现ocp原则
//ocp原则：开闭原则，对扩展开放，对修改关闭

type PET interface {
	eat()
	run()
}

type DOG struct {
	name string
}

type CAT struct {
	name string
}

//DOG实现PET接口
func (dog DOG) eat() {
	fmt.Println(dog.name, "吃汉堡...")
}

func (dog DOG) run() {
	fmt.Println(dog.name, "跑操场...")
}

//CAT实现PET接口
func (cat CAT) eat() {
	fmt.Println(cat.name, "吃薯条...")
}

func (cat CAT) run() {
	fmt.Println(cat.name, "跑墙头...")
}

type PERSON struct {
	name string
}

func (person PERSON) care(pet PET) {
	pet.eat()
	pet.run()
}

func main6() {
	dog := DOG {
		name : "旺财",
	}
	cat := CAT {
		name : "猫咪",
	}

	person := PERSON {
		name : "翠翠",
	}

	person.care(dog)
	person.care(cat)	
}