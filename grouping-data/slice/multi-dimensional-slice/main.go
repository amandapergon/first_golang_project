package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "chocolate", "martini"}
	mb := []string{"Miss", "Moneypenny", "strawberry", "bubblegum"}

	fmt.Println(jb)
	fmt.Println(mb)

	//it's a slice os a slice of tring

	md := [][]string{jb, mb}

	fmt.Println(md)

	fmt.Println(md[1:])
}
