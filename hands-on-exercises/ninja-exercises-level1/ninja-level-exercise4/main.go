package main

import (
	"fmt"
)

type ninito int

var x ninito

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)
}
