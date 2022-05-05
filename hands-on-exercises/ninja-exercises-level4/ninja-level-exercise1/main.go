package main

import "fmt"

func main() {
	ar := [5]int{23, 13, 78, 29, 100}
	fmt.Println(ar)

	for i, v := range ar {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", ar)

}
