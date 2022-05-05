package main

import (
	"fmt"
)

//back ticks `` (raw string literal) will include any return, spaces, whatever you want
//I can reassign a new value to a variable, but you can't change a string
func main() {
	s := "Hello, Playground"
	s = "Hello, Amanda"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Println(i, v)
	}
}
