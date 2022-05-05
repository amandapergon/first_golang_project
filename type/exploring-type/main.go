package main

import (
	"fmt"
)

//this is a STATIC PROGRAMMING LANGUAGE, not DYNAMIC.
//a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
//that VARIABLE can only hold VALUES of that TYPE

//DECLARED the VARIABLE with the IDENTIFIER "y"
//is of type int
//and I'm ASSIGNING the value "42""
var y int = 42

//DECLARED the VARIABLE with the IDENTIFIER "z"
//is of type string
//and I ASSIGN the value "Shaken, not stirred"
var z string = "Shaken, not stirred"
var a string = `James said, "Shaken, nor stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
