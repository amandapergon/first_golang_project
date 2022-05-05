package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8, 9}

	//fmt.Println(x)

	x = append(x, 77, 88, 99, 114)
	//fmt.Println(x)

	y := []int{234, 587, 998, 115}
	x = append(x, y...)
	fmt.Println(y)
	fmt.Println(x)

}
