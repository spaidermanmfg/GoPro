package main

import (
	"fmt"
	"time"
)

/* 
Go	定时器
实现定时操作，内部通过channel来实现
*/

func main() {
	//NewTimer
	timer := time.NewTimer(time.Second * 2)//两毫秒
	fmt.Println("Now Time: ", time.Now())
	t1 := <-timer.C//阻塞，实现等待
	fmt.Println("t1: ",t1)	


	//After
	<-time.After(time.Second * 4)
	fmt.Println("....",time.Now())

	//Stop,停止计时器
	timer = time.NewTimer(time.Second)
	go func() {
		<-timer.C
		fmt.Println("func...")
	}()

	s := timer.Stop()
	if s {
		fmt.Println("stop...")
	}
	time.Sleep(time.Second * 3)
	fmt.Println("end...")


	//Reset，重置NewTimer
	timer = time.NewTimer(time.Second * 5)
	timer.Reset(time.Second * 1)
	t1 = <-timer.C
	fmt.Println(t1)
}