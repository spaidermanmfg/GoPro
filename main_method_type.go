package main
//方法接收者类型

import "fmt"

/* 

*/


type Users struct {
	name string 
}

//值传递
func showUser1(user Users) {
	fmt.Printf("user: %v\n", user)
	user.name = "yuto"
	fmt.Printf("user: %v\n", user)
}


//指针传递
func showUser2(user *Users) {
	fmt.Printf("user: %v\n", *user)
	user.name = "yuto"
	fmt.Printf("user: %v\n", *user)
}

//值传递
func (user Users) showUser3() {
	fmt.Printf("user1: %v\n", user)
	user.name = "yuto"
	fmt.Printf("user2: %v\n", user)
}


//指针传递
func (user *Users) showUser4() {
	fmt.Printf("user3: %v\n", *user)
	user.name = "yuto"
	fmt.Printf("user4: %v\n", *user)
}

func main() {
	//值类型
	p1 := Users {
		name: "mark",
	}
	
	//指针类型
	p2 := &Users {
		name: "mark",
	}
	fmt.Printf("p1: %T\n", p1)
	fmt.Printf("p2: %T\n", p2)

	showUser1(p1)
	showUser2(p2)
	fmt.Printf("r1: %v\n----------\n", p1)
	fmt.Printf("r2: %v\n", *p2)

	p1.showUser3()
	p2.showUser4()
	fmt.Printf("r1: %v\n----------\n", p1)//内容没有修改
	fmt.Printf("r2: %v\n", *p2)//内容被修改
	
}