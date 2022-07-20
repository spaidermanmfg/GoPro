package main

import (
	"fmt"
)

//方法

type Number int



func (n *Number) sayHi() {
	*n *= 2
	fmt.Printf("Hello %v Cat\n", *n)
}

func main() {
	n := Number(5)
	n.sayHi()
	fmt.Printf("Hi %v Cat\n", n)
}