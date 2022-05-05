package main

import (
	"fmt"
)

type ninito int

var x ninito
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)

	y = int(x)

	fmt.Println(x)

	fmt.Printf("%T\n", y)

}
