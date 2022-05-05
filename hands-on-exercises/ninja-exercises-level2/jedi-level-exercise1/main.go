package main

import (
	"fmt"
)

func main() {
	a := 29

	//# adds the 0x in front of the hexadecimal number representation

	fmt.Printf("%d\t\t%b\t%#x\n", a, a, a)
}
