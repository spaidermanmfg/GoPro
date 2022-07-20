package main

import "fmt"

func array1() {
	//array declaration
	var a1 [3]int
	var a2 [5]string
	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", a2)
}

func array2() {
	//初始化数组,指定长度
	var a = [3]int{2, 4, 5}
	var b = [2]string{"hello", "world"}
	var c = [2]bool{true, false}
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
}

func array3() {
	//初始化数组，省略长度
	var a = [...]int{23, 435, 77, 786, 4}
	var b = [...]string{"23", "345dsf", "asdg"}
	var c = [...]bool{true, false, false, true}
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
}

func array4() {
	//指定索引来进行初始化数组,没有指定的会被0填充
	var a = [...]int{0: 2, 2: 4, 6: 5, 7: 8}
	fmt.Printf("a: %v\n", a)
}

func array5() {
	//通过索引来修改数组元素
	var a = [...]int{3, 5, 6, 6, 8}
	a[0] = 888
	fmt.Printf("a: %v\n", a)
}

func array6() {
	//通过索引来访问数组元素
	var a [3]int
	a[0] = 100
	a[1] = 200
	a[2] = 400
	fmt.Printf("a: %v\n", a)
}

func array7() {
	//根据数组长度遍历数组,len(a)
	var a = [...]int{1, 2, 3, 4, 5, 6, 67, 43}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%v] = %v\n", i, a[i])
	}
}

func array8() {
	//通过for range来遍历数组
	var b = [...]int{1, 2, 3, 4, 5, 6, 7, 6}
	for i, v := range b {
		fmt.Printf("b[%v] = %v\n", i, v)
	}
}

func maincc() {
	array1()
	array2()
	array3()
	array4()
	array5()
	array6()
	array7()
	array8()
}
