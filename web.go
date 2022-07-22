package main

//Web应用程序
import (
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("hello world!")
	if _, err := writer.Write(message); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", viewHandler)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil { //启动web服务器
		log.Fatal(err)
	}
}
