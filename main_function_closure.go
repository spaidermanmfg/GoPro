package main
import "fmt"

/* 
闭包
*/

//闭包示例1
func add() func(y int) int {
	var x int
	return func(y int) int {
		x = x + y
		return x
	}
}

//闭包示例2
func suffix(suffix string) func(string) string {
	return func(name string) string {
		return name + suffix
	}
}

//闭包示例3
func callfun(base int) (func(int) int, func(int) int){
	add := func(x int) int {
		return base + x
	}

	sub := func(x int) int {
		return base - x
	}
	return add, sub
}


func mainkk() {
	//在ff的生命周期内，x一直有效
	ff := add()
	re := ff(10)
	fmt.Printf("re: %v\n", re)
	re = ff(30)
	fmt.Printf("re: %v\n---------------\n", re)
	//重新调用建立新的生命周期
	ff = add()
	re = ff(50)
	fmt.Printf("re: %v\n", re)

	//闭包示例2调用
	s := suffix(".pdf")
	r := s("golang")
	fmt.Printf("r: %v\n", r)

	r = s("java")
	fmt.Printf("r: %v\n", r)

	s = suffix(".doc")
	r = s("golang")
	fmt.Printf("r: %v\n", r)

	//闭包示例3调用
	ad, su := callfun(400)
	r1 := ad(400)
	r2 := su(200)
	fmt.Printf("r1: %v\nr2: %v\n", r1, r2)
}