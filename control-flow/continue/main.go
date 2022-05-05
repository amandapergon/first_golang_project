package main

import "fmt"

//break: STOP
//continue: do nothing if this condition is true but go continue the loop

func main() {
	x := 0
	for {
		x++
		if x > 20 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
	fmt.Println("done.")
}
