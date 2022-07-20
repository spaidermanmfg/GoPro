//结构体指针
package main

import "fmt"

type Animalls struct {
	cat     string
	dog     string
	numbers int
}

func struct3() {
	var pepsi Animalls
	pepsi = Animalls{
		cat:     "cat",
		dog:     "dog",
		numbers: 2,
	}
	//结构体指针
	var p *Animalls
	p = &pepsi
	fmt.Printf("pepsi: %v\n", pepsi)
	fmt.Printf("p: %p\n", p)
	fmt.Printf("p.cat: %v\n", *p)

	//使用new关键字创建结构体指针
	var q = new(Animalls)
	q = &pepsi
	fmt.Printf("q: %p\n", q)
	fmt.Printf("q: %v\n", *q)

}

func mainsq() {
	struct3()
}
