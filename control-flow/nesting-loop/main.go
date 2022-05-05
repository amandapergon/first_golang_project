package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 3; i++ {
		fmt.Printf("The outer loop: %d\t\n", i)
		for j := 0; j <= 5; j++ {
			fmt.Printf("\t\tThe inner loop: %d\n", j)
		}
	}
}
