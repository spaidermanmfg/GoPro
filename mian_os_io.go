package main

import (
	"os"
	"fmt"
)
/* 
标准库os   go/src/os
文件目录操作io   go/src/io

*/

//create file
func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Println("err: ", err)
	}else {
		fmt.Println("f.Name(): ", f.Name())
	}
}

//create dir
func makeDir() {
	//创建单个目录
	err := os.Mkdir("globel", os.ModePerm)//目录名，目录权限
	if err != nil {
		fmt.Println("err: ", err)
	}


	//创建多个目录
	err2 := os.MkdirAll("a/b/c", os.ModePerm)
	if err2 != nil {
		fmt.Println("err2: ",err2)
	}
}

//删除文件和目录
func delFile() {
	//删除文件
	err := os.Remove("a.txt")
	if err != nil {
		fmt.Println("romove err: ", err)
	}

	//删除目录
	err3 := os.Remove("globel")
	if err != nil {
		fmt.Println("romove err: ", err3)
	}

	err2 := os.RemoveAll("a")
	if err2 !=nil {
		fmt.Println("removeall: ", err2)
	}

}


//获取当前工作目录
func getNowDir() {
	dir, _ := os.Getwd()
	fmt.Println(dir)

	//修改目录
	// os.Chdir("/Users/")
	// dir, _ = os.Getwd()
	// fmt.Println(dir)


	//获取当前临时目录
	re := os.TempDir()
	fmt.Println("s: ", re)
}

func main() {
	createFile()
	makeDir()
	delFile()
	getNowDir()
}