package main
import "fmt"
import "time"
//import "sync"
import "sync/atomic"

/* 
原子变量


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
} 




