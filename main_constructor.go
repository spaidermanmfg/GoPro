package main

import "fmt"

//构造函数

type Human struct {
	name string
	age int
}

func NewHuman(name string, age int) (*Human, error) {
	if name == "" {
		return nil, fmt.Errorf("name is empty")
	}
	if age <= 0 {
		return nil, fmt.Errorf("age is %d", age)
	}
	return &Human{name : name, age : age,}, nil
}

func main() {
	res, err := NewHuman("", 12)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(*res)
	}
}
