package main

import (
	"fmt"
	"time"
)

/*
Ticker  周期执行

*/
func main() {
	ticker := time.NewTicker(time.Second * 3) //每隔3秒钟
	/*
		counter := 1
		for _ = range ticker.C {
			fmt.Println("codeing...")
			counter++//设置停止条件
			fmt.Println(counter)s
			if counter > 10 {
				ticker.Stop()
				break
			}
		} */

	chanINT := make(chan int)
	go func() {
		for _ = range ticker.C {
			select {
			case chanINT <- 1:
			case chanINT <- 2:
			case chanINT <- 3:
			}
		}
	}()

	sum := 0
	for v := range chanINT {
		fmt.Println("copy ", v)
		sum += v
		if sum >= 10 {
			break
		}
	}

}
