package main

import (
	"fmt"
	"sync"
)
/*
等待组
*/

var wp sync.WaitGroup

func showInt(i int) {
	defer wp.Done()
	fmt.Printf("i = %v\n", i)
}


func main22() {
	for i := 0; i < 10; i++ {
		go showInt(i)
		wp.Add(1)
	}

	wp.Wait()
	//主协程
	fmt.Println("end...")
}