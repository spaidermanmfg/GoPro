package main
//break语句
import (
	"fmt"
)

func r1() {
	//在for循环中使用break语句
	i := 0
	for i <= 10 {
		fmt.Printf("%dss\n", i)
		i++
		if i >= 6 {
			break
		}
	}
	fmt.Println()
}

func r2() {
	//在switch中使用break语句
	//switch语句中break和fallthrough配合使用才有意义
	i := 2
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
		break
		//fallthrough
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	}
}

func r3() {
	//break标签
	MYSQLUU:

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d\n", i)
		if i >= 6 {
			break MYSQLUU
		}
	}
	fmt.Println("break.....")
}

func mainy() {
	r1()	
	r2()
	r3()
}