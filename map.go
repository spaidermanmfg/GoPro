package main

import(
	"log"
	"bufio"
	"os"
	"fmt"
)

func main() {
	fileName := "/Users/mac/Desktop/name.txt"
	lines, err := GetStrings(fileName)
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
	//CreateMap()
	for key, value := range counts {
		fmt.Printf("%v : %v\n", key, value)
	}
}


func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, err
}


func FuncCount(lines []string){
	var names []string
	var counts []int 
	for _, value := range lines {
		flag := false
		for i, name := range names{
			if name == value {
				counts[i]++
				flag = true
			}
		} 
		if flag == false {
			names = append(names, value)
			counts = append(counts, 1)
		}
	}
	for i, name := range names {
		fmt.Printf("%s : %d\n", name, counts[i])
	}
}


func CreateMap(){
	//var ranks map[string]int
	ranks := make(map[string]int)
	ranks["mark"] = 1
	ranks["alex"] = 2
	ranks["tony"] = 3
	fmt.Println(ranks["alex"])

	//映射字面量
	height := map[string]int{
		"a":1,
		"b":2,
		"c":3,
	}
	var value int
	var ok bool
	value, ok = height["a"]
	fmt.Println(value, ok)
	value, ok = height["r"]
	fmt.Println(value, ok)
	delete(height, "a")
	value, ok = height["a"]
	fmt.Println(value, ok)

}