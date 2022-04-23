//接口的值类型接收者和指针类型接收者

package main

import "fmt"

//定义方法
type Pet interface {
	eat(string) string
	eats(string) string
}


type Cat struct {
	name string
}

//实现方法
func (cat Cat) eats(food string) string {
	cat.name = "miao"
	fmt.Printf("%v eat %v\n", cat.name, food)//miao eat fish
	return cat.name
}

//实现方法
func (cat *Cat) eat(food string) string {
	cat.name = "miao"
	fmt.Printf("%v eat %v\n", cat.name, food)
	return cat.name
}



func main() {
	cat := Cat{
		name : "kitty",
	}
	re := cat.eats("fish")
	fmt.Printf("%v\n", re)//miao
	fmt.Printf("%v\n", cat)//kitty

	fmt.Println("-----------------")

	cat1 := &Cat{
		name : "kitty",
	}
	rs := cat1.eat("fish")
	fmt.Printf("%v\n", rs)
	fmt.Printf("%v\n", *cat1)//值被修改
	
}