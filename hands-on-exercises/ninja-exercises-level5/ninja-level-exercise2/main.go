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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	//fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.favIcecreamFlavors {
			fmt.Println(i, val)
		}
	}
}
