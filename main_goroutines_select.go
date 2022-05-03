package main

import (
	"fmt"
	"time"
)

/*
go并发编程之select
类似于switch结构，主要用于出来异步IO操作，当case中channel读写操作能力为非阻塞状态，将会出发相应的动作

select语句中的case语句必须是一个channel操作
select中的default子句总是可以运行的


*/

var chansInt = make(chan int)
var chansStr = make(chan string)

func main() {
	go func() {
		chansInt <- 100
		chansStr <- "globel"
		defer close(chansInt)
		defer close(chansStr)

	}()

	for{
		select {
		case r := <-chansInt:
			fmt.Println("chansInt",r)
		case r := <-chansStr:
			fmt.Println("chansStr",r)
		default:
			fmt.Println("default...")
		}
		time.Sleep(time.Second)
	}
	
}