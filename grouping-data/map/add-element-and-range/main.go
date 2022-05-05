package main

import "fmt"

func main() {
	m := map[string]int{
		"James":  32,
		"Miss":   27,
		"Amanda": 29,
		"Andre":  32,
	}

	m["Dobby"] = 2
	m["Paçoca"] = 1

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
