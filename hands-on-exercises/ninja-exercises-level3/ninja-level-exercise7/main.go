package main

import "fmt"

func main() {
	x := "André"

	if x == "Amanda" {
		fmt.Printf("Hello, %v\n", x)
	} else if x == "André" {
		fmt.Println("Hello, André")
	}
}
