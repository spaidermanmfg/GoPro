package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"log"
	"strconv"
)

func main() {
	k := math.Floor(3.33)
	l := strings.Title("head first go")
	fmt.Println(k)
	fmt.Println(l)

	//用户键盘输入
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin) //设置从键盘获取文本的缓冲读取器
	input, err := reader.ReadString('\n')    //返回用户输入的所有内容
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)//删除字符串前后空白字符
	grade, err := strconv.ParseFloat(input, 64)//将字符串转换为浮点数
	if err != nil {
		log.Fatal(err)//短变量声明中，两个变量其中一个已经声明过，则为第二次短变量声明中为赋值
	}
	var status string
	if grade >= 60 {
		status = "passing"

	}else{
		status = "failing"
	}
	fmt.Printf("A grade of %v is %v", grade, status)
	
}
