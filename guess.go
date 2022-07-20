package main

/*
猜数字游戏
1.生成随机数并保存
2.让玩家猜测
3.最多可以猜测10次
*/

import (
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

func main() {
	seconds := time.Now().Unix()//Unix()将时间转化为整数
	//fmt.Println(seconds)
	rand.Seed(seconds)//播种随机数生成器
	target := rand.Intn(100) + 1//生成1 - 100的随机数
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
    fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)//创建一个bufio.Reader，允许读取键盘输入

	success := false
	for gus := 0 ; gus < 5; gus++ {
		fmt.Printf("You have %v guesses left.\n", 5 - gus)

		fmt.Print("Enter a number: ")
	input, err := reader.ReadString('\n')//读取用户输入，直至键入回车
	if err != nil {
		log.Fatal(err)//异常处理
	}

	input = strings.TrimSpace(input)//删除字符串前后空白字符
	guess, err := strconv.Atoi(input)//将键入内容转化为整数
	if err != nil {
		log.Fatal(err)//异常处理
	}

	fmt.Println(guess)

	if guess > target {
		fmt.Println("Oops. Your guess was High.")
	} else if guess < target {
		fmt.Println("Oops. Your guess was Low.")
	} else {
		fmt.Println("Oops. Your guess was Right.")
		success = true
		break//跳出循环
	}
	}

	if !success {
		fmt.Printf("Sorry, you didn't guess my number. It was: %v \n", target)
	}
	
}