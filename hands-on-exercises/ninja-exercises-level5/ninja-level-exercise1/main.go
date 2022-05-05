package main

import "fmt"

type person struct {
	firstName          string
	lastName           string
	favIcecreamFlavors []string
}

func main() {
	p1 := person{
		firstName: "Amanda",
		lastName:  "Gonçalves",
		favIcecreamFlavors: []string{
			"strawberry",
			"chocolate",
			"vanilla",
		},
	}

	p2 := person{
		firstName: "André",
		lastName:  "Nakazima",
		favIcecreamFlavors: []string{
			"chocolate",
			"doce-de-leite",
			"nutella",
		},
	}

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for i, v := range p1.favIcecreamFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	for i, v := range p2.favIcecreamFlavors {
		fmt.Println(i, v)
	}
}
