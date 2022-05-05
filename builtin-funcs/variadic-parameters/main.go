package main

import "fmt"

// func Println(a ...interface{}) (n int, err error)
// a ...interface{} = what the func accept as arguments:
// - a =
// - ... = unlimited number of them (a variadic parameter)
// - interface{} = it will take a value of any type
// n int, err error = what the func returns:
// - n = number of bytes written
// - err error = any write error it encounters
//When you see that a func has a parameter of this type "...interface{}" this is called a "variadic parameter"
//and it means that the func can take as many values of that type as you want to pass in.

func main() {
	// in here we're catching the returns. This is the syntax to catch a return.
	n, err := fmt.Println("Amanda", 29, true)
	fmt.Println(n)
	fmt.Println(err)

	// the underscore means we ignoring/throwing one of the returns away. It cannot be empty.
	bytes, _ := fmt.Println("André", 32, true)
	fmt.Println(bytes)
}
