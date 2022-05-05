package main

import "fmt"

// if you ever need to declare a variable outside of a function body, use a VAR - it's global scope.
//BEST PRACTICES:
// Limit the scope of your variables and as much as possible try to use a short declaration operator

//DECLARE the variable
//ASSIGN the value 43
//declare & assign = initilization

var y = 43

// DECLARE there is a VARIABLE with the identifier "z"
// and that the variable with the IDENTIFIER "z" is of type int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// ZERO VALUE IS false for booleans, 0 for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels ans maps

var z int

func main() {
	x := 42
	fmt.Println(x)
	//if you ever need to declare a variable inside a function body, use a short declaration operator

	fmt.Println(y)

	foo()
}

func foo() {
	fmt.Println("again:", y)
}
