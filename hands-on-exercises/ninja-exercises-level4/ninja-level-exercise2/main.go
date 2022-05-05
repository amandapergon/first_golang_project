package main

import "fmt"

func main() {
	sl := []int{12, 14, 16, 18, 20, 22, 24, 26, 28, 32}

	for i, v := range sl {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", sl)
}
