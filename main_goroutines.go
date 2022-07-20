package main

import "fmt"
import "time"
import "log"
import "net/http"
import "io/ioutil"

//Go 协程
//Go 协程是由 Go 运行时管理的轻量级线程。
/*
使用 go 关键字开始一个协程，就会启动一个新的协程。
*/

func showMes(mes string) {
	for i := 0; i < 5; i++ {
		fmt.Println("mes:", mes)
	}
	time.Sleep(time.Millisecond * 5) //暂停5毫秒
}

func responseSize(url string) {
	fmt.Println("step1:", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("step2:", url)
	defer response.Body.Close()

	fmt.Println("step3:", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("step4:", len(body))
}

func main9() {
	go showMes("china") //go 关键字开启一个协程
	go showMes("American")
	time.Sleep(time.Millisecond * 2) //暂停2毫秒,
	fmt.Println("main over")         //主函数结束，程序退出

	go responseSize("http://www.baidu.com")
	go responseSize("http://www.qq.com")
	go responseSize("http://www.minfg.top")
	time.Sleep(time.Second * 2) //暂停2秒
}
