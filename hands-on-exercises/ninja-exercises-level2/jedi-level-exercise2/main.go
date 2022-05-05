package main

import (
	"fmt"
)

func main() {
	a := (29 == 29)
	b := (29 <= 30)
	c := (29 >= 28)
	d := (29 != 92)
	e := (29 < 30)
	f := (29 > 28)

	fmt.Println(a, b, c, d, e, f)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
}
