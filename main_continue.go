package main

//continue关键字

import (
	"fmt"
)

func g1() {
	fmt.Println()
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d\n", i)
		} else {
			continue
		}
	}
}

func g2() {
	for i := 0; i <= 10; i++ {
	END:
		for j := 0; j <= 10; j++ {
			fmt.Printf("%d , %d\n", i, j)
			if i == 6 && j == 6 {
				continue END
			}
		}
	}
}

func mainb() {
	//Go中continue关键字只能在for循环中使用
	g1()
}
