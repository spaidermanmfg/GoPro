package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strconv"
)

func main() {
	fileName := "/Users/mac/Desktop/data.txt"
	//numbers, err := GetArrays(fileName)
	numbers, err := GetSlice(fileName)
	if err != nil {
		log.Fatal(err)
	} else{
		var sum float64 = 0.0
		for _, value := range numbers {
			sum += value
		}
		count := float64(len(numbers))
		fmt.Printf("Average: %.2f\n", sum/count)
	}
}


func GetArrays(fileName string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(fileName)////打开数据文件
	if err != nil {
		log.Fatal(err)
		return numbers, err
	}
	var i int = 0
	scanner := bufio.NewScanner(file)////为文件创建一个扫描器
	for scanner.Scan() {////从文件中读取一行
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)//转为float64类型，并存入数组中
		if err != nil {
			log.Fatal(err)
			return numbers, err
		}
		i++
	}
	err = file.Close()//释放资源
	if err != nil {
		return numbers, err
	} 
	
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}

func GetSlice(fileName string) ([]float64, error){
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
} 