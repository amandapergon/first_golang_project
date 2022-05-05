package main

import "fmt"

func main() {
	j := []string{"James", "Bond", "Shaken, not stirred"}
	m := []string{"Miss", "Moneypenny", "Helloooooo, James"}

	md := [][]string{j, m}
	fmt.Println(md)

	for i, v := range md {
		fmt.Println(i, v)
	}

	for _, v := range md {
		for j, v := range v {
			fmt.Printf("\t%v\t%v\n", j, v)
		}
	}
}
