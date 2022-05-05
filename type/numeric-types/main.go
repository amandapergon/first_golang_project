package main

import (
	"fmt"
	"runtime"
)

// the assignment below doesn't work because it overflows int8
//var w int8 = -129
var x int
var y float64
var z int8 = -128

func main() {

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	x = 2
	y = 42.7125
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

}
