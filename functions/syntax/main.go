package main

import "fmt"

func main() {
	basic()
	takesAnArgument("Takes an argument func!")
	s1 := returnFunc("Return func!")
	fmt.Println(s1)
	x, y := multipleReturns("Multiple", "returns")
	fmt.Println(x)
	fmt.Println(y)
}

//func (r receiver) identifier(parameters) (return(s)) { ... }

func basic() {
	fmt.Println("Hello from Basic func!")
}

func takesAnArgument(s string) {
	fmt.Println("Hello from", s)
}

func returnFunc(s string) string {
	return fmt.Sprint("Hello from ", s)
}

func multipleReturns(s string, st string) (string, bool) {
	a := fmt.Sprint(s, " ", st, ` says hello!`)
	b := true
	return a, b
}
