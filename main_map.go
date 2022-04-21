package main
//map

import "fmt"

func map1() {
	//使用map关键字声明map语法
	//var mapName map[keyType]valueType
	//使用make函数声明map
	//var mapName = make(map[keyType]valueType])
	var info map[string]string//声明
	info = make(map[string]string)//初始化
	fmt.Printf("%T\n", info)
	fmt.Println(info)

	//直接初始化
	var john = map[string]string{
		"name":"john",
		"age":"25",
		"job":"skate",
	}
	fmt.Printf("john: %v\n", john)

	//使用make初始化
	yuri := make(map[string]string)
	yuri["name"] = "yuri"
	yuri["age"] = "24"
	yuri["job"] = "doctor"
	fmt.Printf("yuri: %v\n", yuri)
}

func map2() {
	//Access to map elements using a key
	maps := make(map[string]string)
	maps["name"] = "kimi"
	maps["age"] = "41"
	maps["job"] = "driver"

	var key = "job"
	var value = maps[key]
	fmt.Printf("%v : %v\n", key, value)
}

func map3() {
	//determine whether a map contains a certain key
	//value, ok := mapName[key]
	var maps = make(map[string]string)
	maps["name"] = "kimi"
	maps["age"] = "41"
	maps["job"] = "driver"

	var key1 = "age"
	var key2 = "email"

	v, ok := maps[key1]
	fmt.Printf("v : %v, ok : %v\n", v, ok)

	v, ok = maps[key2]
	fmt.Printf("v : %v, ok : %v\n", v, ok)
}

func map4() {
	//ergodic map elements using range loop 
	var maps = make(map[string]string)
	maps["name"] = "kimi"
	maps["age"] = "41"
	maps["job"] = "driver"
	maps["award"] = "gold"

	for key, value := range maps {
		fmt.Printf("%v : %v\n", key, value)
	}
}

func mainpp() {
	map1()
	map2()
	map3()
	map4()	
}