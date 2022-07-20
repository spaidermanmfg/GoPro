package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", height)
	}
	area := width * height
	//fmt.Printf("%.2f liters needed\n", area/10.0)
	return area / 10.0, nil
}

func main() {
	area, err  := paintNeeded(3.4, -5.6)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("area: %0.2f\n", area)
	}	
}