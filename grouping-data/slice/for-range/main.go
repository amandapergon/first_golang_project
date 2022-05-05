package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8, 9}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])

	y := []string{"Amanda", "André", "Dobby", "Paçoca"}

	fmt.Println(y)
	fmt.Println(y[1])

	//for i, v := range x {
	//	fmt.Println(i, v)
	//}

	//for i, v := range y {
	//	fmt.Println(i, v)
	//}

	//for i := 0; i <= 3; i++ {
	//	fmt.Println(i, y[i])
	//}

	//for i := 0; i <= 5; i++ {
	//	fmt.Println(i, x[i])
	//}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
