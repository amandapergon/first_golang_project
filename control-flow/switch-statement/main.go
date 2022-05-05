package main

import "fmt"

//fallthrough:
// transfers control to the first statement of the next case clause in an expression "switch" statement

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print?")
		fallthrough
	case (7 == 9):
		fmt.Println("not true")
		fallthrough
	case (31 == 43):
		fmt.Println("not true2")
		fallthrough
	case (15 == 15):
		fmt.Println("true 15")
		fallthrough
	default:
		fmt.Println("this is default")
	}

	n := "Bond"
	switch n {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("Bond james")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is the default")
	}

	switch n {
	case "Money penny", "Bond", "Do no":
		fmt.Println("miss money or bond or Do no")
	case "M":
		fmt.Println("M")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is the default")
	}

}
