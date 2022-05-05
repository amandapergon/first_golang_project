//every program needs package main
package main

//control flow - the order in which individual statements, instructions or function calls of an imperative program are executed or evaluated
//(1) sequence
//(2) loop; iterative
//(3) conditional

import "fmt"

//every program needs func main
// your program is done when you leave func main. it's OVER

func main() {
	fmt.Println("Hello, Amanda!")
	foo()
	fmt.Println("Hello again, Amanda")

	for i := 0; i < 100; i++ {
		if i%10 == 0 {
			fmt.Println("Hello once more, Amanda!")
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("I'm in bar")
}
