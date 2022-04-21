package main
import "fmt"
//Go数据类型

func tough() {
	//函数类型
}

func maina() {

	//打印数据类型
	var name string = "张三"
	age := 18
	model := true
	p := &age //指针类型
	a := [...]int{1, 2, 3, 4, 5}  //array类型
	b := []int{1, 2, 3, 4, 5}  //slice类型,可以动态添加
	//打印值
	fmt.Printf("name:%s, age:%d, model:%t\n",name,age,model)
	fmt.Printf("p:%p\n",p)
	fmt.Printf("a:%v\n",a)
	fmt.Printf("b:%v\n",b)
	//打印数据类型
	fmt.Printf("-----------------------------------------------------\n")
	fmt.Printf("name:%T, age:%T, model:%T\n",name,age,model)
	fmt.Printf("p:%T\n",p)
	fmt.Printf("a:%T\n",a)
	fmt.Printf("b:%T\n",b)
	fmt.Printf("tough:%T\n", tough)

	
}