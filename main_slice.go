package main
//切片

import "fmt"

func slice1() {
	//声明切片,不需要指定长度
	//var slice []int
	var a []int
	var b []string
	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", b, b)
}

func slice2() {
	//通过make函数声明切片,cap()返回切片的容量,len()返回切片的长度
	var a = make([]int, 3)
	fmt.Printf("%T %v len(a)=%v cap(a)=%v\n", a, a, len(a), cap(a))
}

func slice3(){
	//直接初始化切片
	var a = []int{1,2,3,4,5,6}
	fmt.Printf("%v\n", a)
	//切片的切片,取0但不取3
	b := a[0:3]
	fmt.Printf("%v\n", b)
	//取3到结束
	d := a[3:]
	fmt.Printf("%v\n", d)
	//取0到结束
	e := a[:]
	fmt.Printf("%v\n", e)
}

func slice4() {
	//根据索引和长度遍历切片
	var a = []int{1,2,3,4,5,6,7,8,9,0}
	length := len(a)
	for i := 0; i < length; i++ {
		fmt.Printf("a[%v]=%v\n", i, a[i])
	}
}


func slice5() {
	//通过for range遍历切片
	var a = []int{1,2,4,5,6,7,8,9,0}
	for i, v := range a {
		fmt.Printf("a[%v]=%v\n", i, v)
	}
}

//crud,切片的增删改查
//add
func slice6() {
	//使用append()函数向切片中添加元素
	var a = []int{}
	a = append(a, 100)
	a = append(a, 200)
	a = append(a, 300)
	fmt.Printf("a: %v\n", a)
}

//del
func slice7() {
	//删除切片中的元素
	var a = []int{1,2,3,4,5,6,7,8,9,0}
	//删除第一个元素
	a = append(a[:0], a[1:]...)
	fmt.Printf("a: %v\n", a)
	//删除最后一个元素
	a = append(a[:len(a)-1], a[len(a):]...)
	fmt.Printf("a: %v\n", a)
	//删除5
	a = append(a[:3], a[4:]...)
	fmt.Printf("a: %v\n", a)
}

//update
func slice8() {
	//更新切片中的元素
	var a = []int{1,2,3,4,5,6,7,8,9}
	a[4] = 1000
	fmt.Printf("a: %v\n", a)
}

//query
func slice9() {
	//查询切片中的元素
	var a = []int{1,2,3,4,5,6,7,8,9,0}
	key := 4
	for i, v := range a {
		if i == key {
			fmt.Printf("a[%v] = %v\n", i, v)
		}
	}
}

//赋值切片
func slice10() {
	//赋值切片会赋值切片的地址，修改其中一个切片，另一个切片也会被修改
	var a = []int{1,2,3,4,5}
	var b = a
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	b[2] = 500
	fmt.Println("------------------")
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
}

//使用copy(目标切片，源切片)函数复制切片
func slice11() {
	var a = []int{1,2,3,4,5,6}
	var b = make([]int, len(a))
	copy(b, a)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)	
}

func mainss() {
	slice1()
	slice2()
	slice3()
	slice4()
	slice5()
	slice6()
	slice7()
	slice8()
	slice9()
	slice10()
	slice11()
}