package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	licenseToKill bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   35,
		},
		licenseToKill: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
			age:   28,
		},
		licenseToKill: true,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)

	fmt.Println(sa1.first, sa1.last, sa1.age)
	fmt.Println(sa2.first, sa2.last, sa2.age)
}
