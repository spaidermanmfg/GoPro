package main

import (
	"fmt"
	"runtime"
)
/*
runtime包

runtime.Goexit()//退出当前协程
runtime.Gosched()//让出CPU时间片
runtime.NumCPU()//获取CPU核心数

*/

func showInfos(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %s\n", msg)
		if i == 5 {
			runtime.Goexit()//退出当前goroutine
		}
	}
}


func main() {
	go showInfos("Globel WAR")
	fmt.Println("runtime.NumCPU(): ", runtime.NumCPU())//获取CPU核心数
	runtime.GOMAXPROCS(1)//设置最大CPU数
	for i := 0; i < 3; i++ {
		runtime.Gosched()//让出CPU时间片
		fmt.Println("golang")
	}

	fmt.Println("end...")
	
}