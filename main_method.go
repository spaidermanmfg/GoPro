//Go方法
package main

import "fmt"

/* 
方法就是含有接收者的函数
语法格式：
func (receiverType receiverVariable) functionName(parameterType) returnType {
	function body
}

receiverType：接收者类型，可以是值类型或者指针类型
receiverVariable：接收者变量
functionName：方法名
parameterType：参数类型
returnType：返回值类型

1.接收者的类型可以是struct，slice，map，channel，function，pointer，interface等类型

方法的调用：
1.按照接收者的类型调用方法
2.按照接收者的值调用方法
3.按照接收者的指针调用方法
*/


type User struct {
	name string
}

//(user User) 接收者receive
func (user User) eat() {
	fmt.Printf("%v,eat...\n", user.name)

}

func (user User) sleep() {
	fmt.Printf("%v,sleep...\n", user.name)
}

type Customer struct {
	name string
	pwd string
}

func (customer Customer) login(name string, pwd string) string {
	fmt.Printf("%v, %v\n", customer.name, customer.pwd)
	if name == customer.name && pwd == customer.pwd {
		return "login success..."
	}else{
		return "login failed..."
	}

}
	

func main() {
	tmp := User{
		name: "mark",
	}
	tmp.eat()
	tmp.sleep()


	cus := Customer{
		name: "fard",
		pwd: "6767",
	}
	re := cus.login("fard", "6767")
	fmt.Printf("%v\n", re)
}