package main

import (
	"fmt"
)

//use short declaration operator as much as possible
//use var for:
//zero value
//package scope

var y string
var z int

func main() {
	//DECLARE  a variable to be of a certain TYPE
	// and then ASSIGN a VALUE of that type to the variable

	//zero value is the initial, default value before we ASSIGN one ourselves

	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)

	y = "Hello, Amanda"

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println("printing the value of z", z, "ending")
	fmt.Printf("%T\n", z)

	z = 29

	fmt.Println(z)
	fmt.Printf("%T\n", z)
}

//false for booleans
//0 for integers
//0.0 for floats
//"" for strings
//nil for
//pointers
//functions
//interfaces
//slices
//channels
//maps
