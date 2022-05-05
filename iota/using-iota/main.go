package main

import "fmt"

const (
	//this first iota is equal to zero
	//_ = using _ so we do not need to use the variable
	_ = iota
	//kb = 1024, which is the same as 1 kilobyte, it's one SHIFTED OVER 10x
	kb = 1 << (iota * 10)
	mg = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mg, mg)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
