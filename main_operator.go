package main
import(
	"fmt"
)

//算术运算符

func mainn() {
	var a int = 100
	var b int = 5

	//算术运算符
	fmt.Printf("a + b = %d\n", a + b)
	fmt.Printf("a - b = %d\n", a - b)
	fmt.Printf("a * b = %d\n", a * b)
	fmt.Printf("a / b = %d\n", a / b)
	fmt.Printf("a mod b = %d\n", a % b)

	//Go语言中++、--不能用作表达式中
	a++
	fmt.Printf("a++ = %d\n", a)
	a--
	fmt.Printf("a-- = %d\n\n", a)

	//关系运算符
	judge := a == b
	fmt.Printf("a == b 为 %t\n", judge)
	judge = a != b
	fmt.Printf("a != b 为 %t\n", judge)
	judge = a > b
	fmt.Printf("a > b 为 %t\n", judge)
	judge = a < b
	fmt.Printf("a < b 为 %t\n", judge)
	judge = a >= b
	fmt.Printf("a >= b 为 %t\n", judge)
	judge = a <= b
	fmt.Printf("a <= b 为 %t\n\n", judge)

	//逻辑运算符
	flag1 := true
	flag2 := false

	judge = flag1 && flag2
	fmt.Printf("flag1 && flag2 为 %t\n", judge)
	judge = flag1 || flag2
	fmt.Printf("flag1 || flag2 为 %t\n", judge)
	judge = !flag1
	fmt.Printf("!flag1 为 %t\n\n", judge)


	//位运算符
	var c int = 4//0100
	var d int = 8//1000

	fmt.Printf("c = %b \nd = %b\n", c, d)
	result := c & d
	fmt.Printf("c & d = %b\n", result)//两个都为1的位置为1，否则为0
	result =  c | d
	fmt.Printf("c | d = %b\n", result)//有一个位置为1，则为1，否则为0
	result = c ^ d
	fmt.Printf("c ^ d = %b\n", result)//两个位置不同为1，否则为0
	result = c << 2
	fmt.Printf("c << 2 = %b\n", result)//左移2位
	result = c >> 2
	fmt.Printf("c >> 2 = %b\n\n", result)//右移2位

	//赋值运算符
	a += 100
	fmt.Printf("a += 5 = %d\n", a)
}