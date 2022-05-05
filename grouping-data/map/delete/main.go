package main

import "fmt"

func main() {
	m := map[string]int{
		"James":  32,
		"Miss":   27,
		"Amanda": 29,
		"Andre":  32,
	}

	fmt.Println(m)

	delete(m, "Miss")
	fmt.Println(m)

	if v, ok := m["James"]; ok {
		fmt.Println("value:", v)
		delete(m, "James")
	}

	fmt.Println(m)
}
