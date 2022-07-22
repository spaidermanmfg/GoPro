package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

//检索网页
func main() {
	myChannel := make(chan Page)
	urls := []string{
		"https://www.iguopin.com",
		"https://www.bilibili.com",
		"http://zhaopin.csg.cn/#/maintain/index/page",
	}

	for _, url := range urls {
		go responseSize1(url, myChannel)
	}

	for i := 0; i < len(urls); i++ {
		page := <-myChannel
		fmt.Printf("%v : %v\n", page.URL, page.Size)
	}
}

func responseSize1(url string, channel chan Page) {
	//fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)} //向channel中发送数据
	//fmt.Println(len(body)) //将byte类型转化为string
}
