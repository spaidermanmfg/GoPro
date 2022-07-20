package main

import "fmt"

func v1() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d\n", i)
		} else {
			goto END
		}
	}
END:
	fmt.Println("ENDINGs....")
}

func v2() {
	//跳出双层循环
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
			fmt.Printf("%d , %d\n", i, j)
			if i == 6 && j == 6 {
				goto END
			}
		}
	}
END:
	fmt.Println("ENDING...")
}

func maini() {
	//goto关键字,强制跳转
	v1()
	v2()
}
