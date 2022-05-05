package main

import "fmt"

func main() {
	p1 := struct {
		first, last string
		age         int
	}{
		first: "Amanda",
		last:  "Gonçalves",
		age:   29,
	}

	fmt.Println(p1)

}
