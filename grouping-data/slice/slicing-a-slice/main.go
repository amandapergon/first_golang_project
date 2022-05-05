package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8, 9}

	fmt.Println(x[1:])
	//[5 6 7 8 9]
	fmt.Println(x[4:])
	//[8 9]
	fmt.Println(x[2:4])
	//[6 7]
	fmt.Println(x[:2])
	//[4 5]
	fmt.Println(x[2:])
	//[6 7 8 9]

}
