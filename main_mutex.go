package main


import (
	"fmt"
	"sync"
	"time"
)

/*
通过互斥锁实现线程安全


*/

var k int = 100

var wc sync.WaitGroup

var lock sync.Mutex


func aDD() {
	defer wc.Done()
	lock.Lock()//加锁
	k += 1
	fmt.Println("i++ = ", k)
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()//解锁
}

func sUB() {
	lock.Lock()//加锁时其他进程不能进来，只有当程序执行完成解锁后其他线程才能加入进来
	time.Sleep(time.Millisecond * 2)
	defer wc.Done()
	k -= 1
	fmt.Println("i-- = ", k)
	lock.Unlock()
}

func main66() {
	for i := 0; i < 100; i++ {
		wc.Add(1)
		aDD()
		wc.Add(1)
		sUB()
	}

	wc.Wait()

	fmt.Println("end...", k)
	
}