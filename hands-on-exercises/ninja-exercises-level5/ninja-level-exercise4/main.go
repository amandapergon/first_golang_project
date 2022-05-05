package main

import "fmt"

func main() {
	p1 := struct {
		name, occupation string
		age              int
		favFood          []string
		favCharacters    map[string]string
	}{
		name:       "Amanda",
		age:        29,
		occupation: "developer",
		favFood: []string{
			"pasta",
			"hotdog",
			"chocolate",
		},
		favCharacters: map[string]string{
			"The office":                "Dwight",
			"Friends":                   "Chandler",
			"What we do in the Shadows": "Nandor",
		},
	}

	p2 := struct {
		name, occupation string
		age              int
		favFood          []string
		favCharacters    map[string]string
	}{
		name:       "André",
		age:        29,
		occupation: "dba",
		favFood: []string{
			"hamburger",
			"pizza",
			"eggs",
		},
		favCharacters: map[string]string{
			"The office":                "Dwight",
			"Friends":                   "Joey",
			"What we do in the Shadows": "Colin Robinson",
		},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	for i, v := range p1.favFood {
		fmt.Println(i, v)
	}

	for i, v := range p2.favFood {
		fmt.Println(i, v)
	}

	for k, v := range p1.favCharacters {
		fmt.Println(k, v)
	}

	for k, v := range p2.favCharacters {
		fmt.Println(k, v)
	}
}
