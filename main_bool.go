package main

import "fmt"

func mainz() {
	age := 18
	sample := age >= 18
	fmt.Println(sample)
	if sample {
		fmt.Println("成年人")
	} else {
		fmt.Println("未成年人")
	}
	count := 30
	for i := 0; i < count; i++ {
		if i >= age {
			fmt.Println(i, "岁已经是成年人了")
		} else {
			fmt.Println(i, "岁还是未成年人")
		}
	}

}
