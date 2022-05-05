package main

import "fmt"

// if you ever need to declare a variable outside of a function body, use a VAR
// Limit the scope of your variables and as much as possible try to use a short declaration operator

func main() {
	//short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	//if you ever need to declare a variable inside a function body, use a short declaration operator
	x := 45
	fmt.Println(x)
	x = 90
	fmt.Println(x)
	y := 50 * 2
	fmt.Println(y)
	z := "Bond, James"
	fmt.Println(z)
}
