package main

import "fmt"

//a slice is made on the top of an array. when you slice a slice, go has to
//create a new array and throw that one out

//using make declares the size of the slice, so go already knows the size of the array it has to create

//make([]type, lenth, capacity)

//capacity: spots in the underlying array to use

func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 42
	x[9] = 999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 2222)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 4563)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//as I appended more than the capacity stablished, we throw out the old array,
	//copy all its values and create a new array with double the capacity

	x = append(x, 1)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
