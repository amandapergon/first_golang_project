package main

import "fmt"

func main() {
	switch {
	case (2 == 3):
		fmt.Println("This shouldn't print.")
	case (2 > 3):
		fmt.Println("This shouldn't print either.")
	case (2 == 2):
		fmt.Println("This must print.")
		fallthrough
	case (3 > 2):
		fmt.Println("This sould print too, as I used fallthrough")
	}
}
