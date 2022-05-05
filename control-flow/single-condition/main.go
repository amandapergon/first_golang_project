package main

import (
	"fmt"
)

func main() {

	//this is how we do a while statement in go
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("Done")
}
