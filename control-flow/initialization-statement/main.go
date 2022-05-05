package main

import "fmt"

// x is limited to the if scope

func main() {
	if x := 42; x == 42 {
		fmt.Println("001")
	}
	//the below won't run:
	//fmt.Println(x)
}
