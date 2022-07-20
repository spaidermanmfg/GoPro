package main

import (
	"fmt"
	"strings"
)

func f1() {
	//判断奇偶数
	var num int
	fmt.Print("请输入一个数字：")
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Printf("数字%d为偶数\n", num)
	} else {
		fmt.Printf("数字%d为奇数\n", num)
	}
}

func f2() {
	//判断成绩等级
	var score int
	fmt.Print("请输入你的成绩：")
	fmt.Scanln(&score)
	if score < 60 {
		fmt.Println("不及格")
	} else if score >= 60 && score <= 70 {
		fmt.Println("及格")
	} else if score >= 70 && score <= 80 {
		fmt.Println("良好")
	} else if score >= 80 && score <= 90 {
		fmt.Println("优秀")
	} else if score >= 90 && score < 100 {
		fmt.Println("牛逼")
	} else if score == 100 {
		fmt.Println("满分老铁")
	}
}

func f3() {
	//判断星期几
	//Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
	var week string
	fmt.Print("请输入星期：")
	fmt.Scanln(&week)
	if strings.HasPrefix(week, "M") {
		fmt.Println("Today is Monday.")
	} else if strings.HasPrefix(week, "T") {
		fmt.Printf("请输入第二个字母：")
		fmt.Scanln(&week)
		if strings.Contains(week, "u") {
			fmt.Println("Today is Tuesday.")
		} else if strings.Contains(week, "h") {
			fmt.Println("Today is Thursday.")
		}
	} else if strings.HasPrefix(week, "W") {
		fmt.Println("Today is Wednesday.")
	} else if strings.HasPrefix(week, "F") {
		fmt.Println("Today is Friday.")
	} else if strings.HasPrefix(week, "S") {
		fmt.Printf("请输入第二个字母：")
		fmt.Scanln(&week)
		if strings.Contains(week, "a") {
			fmt.Println("Today is Saturday.")
		} else if strings.Contains(week, "u") {
			fmt.Println("Today is Sunday.")
		}
	}
}

func f4() {
	//switch语句,Go语言中没有break语句
	var grade string
	fmt.Print("请输入你的等级：")
	fmt.Scanln(&grade)

	switch grade {
	//单条件匹配
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	case "C":
		fmt.Println("一般")
	//多条件匹配
	case "D", "E", "F", "G", "H":
		fmt.Println("不及格")
	default:
		fmt.Println("还行")
	}

	//case条件可以使用表达式
	var score int
	fmt.Print("请输入你的成绩：")
	fmt.Scanln(&score)
	switch {
	case score < 60:
		fmt.Println("不及格")
	case score >= 60 && score <= 70:
		fmt.Println("及格")
		fallthrough //穿透，跳过下一个case
	case score >= 70 && score <= 80:
		fmt.Println("良好")
	case score >= 80 && score <= 90:
		fmt.Println("优秀")
	case score >= 90 && score < 100:
		fmt.Println("牛逼")
	case score == 100:
		fmt.Println("满分老铁")
	default:
		fmt.Println("还行")
	}
}

//Go流程控制
func mainj() {
	fmt.Println("Go流程控制")

	//if语句,输入变量
	var age int
	fmt.Print("请输入一你的年龄：")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("成年人, 年龄为", age)
	} else {
		fmt.Println("未成年人， 年龄为", age)
	}

	//初始变量可以声明在布尔表达式中，但作用域仅在if语句中可以使用
	//不能使用0或非0来表示真假
	if age := 9; age > 18 {
		fmt.Println("成年人, 年龄为", age)
	} else {
		fmt.Println("未成年人， 年龄为", age)
	}

	f1()
	f2()
	f3()
	f4()
}
