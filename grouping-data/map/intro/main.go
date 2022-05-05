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
	fmt.Println(m["Amanda"])
	fmt.Println(m["Dobby"])

	m["Dobby"] = 2
	m["Paçoca"] = 1

	fmt.Println(m)

	//v, ok := m["Paçoca"]

	//fmt.Println(v)
	//fmt.Println(ok)

	//if v, ok := m["Miss"]; ok {
	//	fmt.Println("This is the if print", v)
	//}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
