package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

//the values assigned to the variables by default are called zero values
