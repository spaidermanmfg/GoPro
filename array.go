package main

import (
	"fmt"
)


func main() {
	notes := [5]string{"go", "java", "python", "rust", "c++"}
	for i := 0; i < len(notes); i++ {
		fmt.Println(notes[i])
	}

	for index, value := range notes {
		fmt.Printf("%v : %v\n", index, value)
	}

	numbers := [4]float64{33.3, 55.7, 66.7, 22.1}
	var sum float64 = 0.0
	for index, value := range numbers {
		fmt.Printf("%v : %v\n", index, value)
		sum += value
	}
	//fmt.Println(sum)
	count := float64(len(numbers))
	fmt.Printf("Average: %.2f\n", sum/count)

}