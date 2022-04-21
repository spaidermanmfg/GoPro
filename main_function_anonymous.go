package main
import "fmt"

/* 
Go 匿名函数
Go 匿名函数是一个简单的语法糖，它可以让我们在没有命名的情况下定义一个函数。
匿名函数的语法与其他语言的匿名函数类似，但是不需要使用 func 关键字。
func (paramlist)(return)

*/

func maindd() {

	//匿名函数,可以嵌套在main函数里
	max := func (a int, b int) int {
		if a > b {
			return a
		}else{
			return b
		}
	}
	re := max(5, 6)
	fmt.Printf("re: %v\n", re)


	//自己调用自己
	sumself := func (a int, b int) int {
		return a + b
	}(6, 7)

	fmt.Printf("sumself: %v\n", sumself)
}