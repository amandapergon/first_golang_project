package main

import "fmt"

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	//%d = prints out the decimal
	//\t\t = prints out a sequence of 2 tabs
	//%b = prints out the binary correspondant to the value
	//\n = breaks a line

	y := x << 1

	//creates number 8

	fmt.Printf("%d\t\t%b\n", y, y)
}
