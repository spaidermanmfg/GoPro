//结构体作为参数传递
package main

import "fmt"


/*
1.直接传递
是一个拷贝，在函数内部不会修改外部结构体的内容
2.使用指针传递
传递结构体指针，在函数内部可以修改外部结构体的内容 
*/


type Book struct {
	title string
	author string
	page int
	bookid int
}

//值传递拷贝了一份副本
func showBook(book Book) {
	book.title = "《old man and fish》"
	book.author = "海明威"
	book.page = 117
	book.bookid = 90
	fmt.Printf("book1: %v\n", book)
}

func showBooks(book *Book) {
	book.title = "《old man and fish》"
	book.author = "海明威"
	book.page = 117
	book.bookid = 90
	fmt.Printf("book1: %v\n", book)
}

func mainpa() {
	book := Book {
		"《Samll King》",
		"Edition",
		222,
		76,
	}
	//没有修改
	fmt.Printf("book2: %v\n", book)
	showBook(book)
	fmt.Printf("book3: %v\n\n", book)

	//值被修改
	var per *Book = &book//per := &book
	fmt.Printf("book4: %v\n", book)
	showBooks(per)
	fmt.Printf("book5: %v\n", per)
	
}