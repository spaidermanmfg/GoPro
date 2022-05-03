package main

import (
	"fmt"
)

/* 
channel的遍历，通过for循环和for range


*/

var y = make(chan int)

func main() {

	go func() {
		for i := 0; i < 5; i++ {
			y <- i
		}
		close(y)
	}()

	// r := <-y
	// fmt.Println(r)
	// r = <-y
	// fmt.Println(r)

	/* //1
	for i := 0; i < 5; i++ {
		rs := <-y
		fmt.Println(rs)
	}

	//2
	for v := range y {
		fmt.Println(v)
	}
 */
	//3
	for {
		v, ok := <-y
		if ok {
			fmt.Println(v)
		}else{
			break
		}
	}
}