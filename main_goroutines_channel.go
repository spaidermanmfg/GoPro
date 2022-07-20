package main

import (
	"fmt"
	"math/rand"
	"time"
)

//channel管道
/*
通道被设计用于在goroutines之间传递数据或资源。
声明通道需要指定数据类型
数据在通道上传递，在指定时间内只有一个goroutine可以接收或发送数据,因此不会发生冲突。
管道分为有缓存和无缓存两种，无缓存的用于同步通信，有缓存的用于异步通信。
使用chan关键字声明。
声明语法：
var chanName chan Type   //无缓存管道
var chanName chan Type = make(chan Type, capacity)   //有缓存管道，capacity为缓存大小，默认为0
使用 <-  运算符将数据发送到通道中，通道在左边表示发送数据，通道在右边表示接收数据。
*/

//创建一个int类型管道
/* var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(100)
	fmt.Printf("send: %v\n", value)
	time.Sleep(time.Second * 3)//睡眠5秒
	values <- value //发送数据
}

func main10() {
	defer close(values)
	go send()
	fmt.Println("wait...")
	value := <- values  //接收数据
	fmt.Printf("receive: %v\n", value)
	fmt.Println("end...")
}
*/
