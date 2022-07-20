package main

import (
	"fmt"
	"os"
	"math"
)

func main() {
	//testSlice()
	fmt.Println(os.Args[1:])
	floatSlice := []float64{22.1, 55.4, 99.2, 44.5}
	fmt.Println(maximum(22.1, 55.4, 99.2, 44.5))
	fmt.Println(maximum(floatSlice...))//使用切片调用
	
	
}


func testSlice() {
	slice := []int{1,2,3,4,5,6}
	fmt.Println(slice, len(slice))
	slice = append(slice, 7, 8, 9)//在末尾追加3个元素
	fmt.Println(slice, len(slice))
}


//可变长参数函数
func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, value := range numbers{
		if value > max {
			max = value
		}
	}
	return max
}