package main
import "fmt"
import "time"
//import "sync"
import "sync/atomic"

/* 
原子操作 atomic
atomic提供的原子操作能确保任一时刻只有一个goroutine对变量进行操作，避免锁操作

atomic常见操作：
	增减
	载入 
	比较并交换cas
	交换
	存储



*/

/* 
var i int = 100
var lock sync.Mutex

func ADD() {
	lock.Lock()
	i++
	lock.Unlock()
}

func SUB() {
	lock.Lock()
	i--
	lock.Unlock()
}

func main() {
	for i := 1; i < 10; i++ {
		go ADD()
		go SUB()
	}

	time.Sleep(time.Second * 2)	
	fmt.Println(i)
	} 
*/

var i int32 = 100

//增减
func ADD(){
	atomic.AddInt32(&i, 1)
}

func SUB(){
	atomic.AddInt32(&i, -1)
}





func main() {
	for i := 1; i < 10; i++ {
		go ADD()
		go SUB()
	}

	time.Sleep(time.Second * 2)	
	fmt.Println(i)

	//载入,读操作
	atomic.LoadInt32(&i)
	fmt.Printf("I: %v\n", i)

	//存储，写操作
	atomic.StoreInt32(&i, 200)
	fmt.Println("store:",i)

	//比较交换，返回bool值
	b := atomic.CompareAndSwapInt32(&i, 100, 200)//与旧值进行比较，如果被修改则停止交换s
	fmt.Printf("b: %v\ni: %v \n", b, i)

} 




