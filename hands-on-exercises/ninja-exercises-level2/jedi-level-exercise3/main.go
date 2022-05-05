package main

import (
	"fmt"
)

//an untyped constant is a constant of a kind, and it gives the compiler
//a little bit more flexibility about how to use that value throughout the program

const (
	a        = "Amanda"
	b string = "Amanda"
)

func main() {
	fmt.Println(a, b)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}
