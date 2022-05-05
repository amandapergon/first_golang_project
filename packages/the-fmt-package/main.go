package main

import (
	"fmt"
)

var y = 42

//	GROUP ONE - general printing to standard out

//func Print(a ...interface{})(n int, err error)
//this one will print it and not give it a line

//func Printf(format string, a ...interface{})(n int, err error)
//this will print and do some formatting. It needs a formatting string and the value you want printed (a ...interface{})

//func Println(a ...interface{})(n int, err error)
//this will print and add a line to the end

// GROUP TWO - printing to a string which you can assisgn to a variable

//func Sprint(a ...interface{}) string

//func Sprintf(format string, a ...interface{}) string

//func Sprintln(a ...interface{}) string

// GROUP THREE - printing to a file or a web server's response

//func Fprint(w io.Writer, a ...interface{})(n int, err error)

//func Fprintf(w io.Writer, format string, a ...interface{})(n int, err error)

//func Fprintln(w io.Writer, a ...interface{})(n int, err error)

//%v is the value in a default format

var greeting string
var name string
var year int

func main() {
	fmt.Println(y)
	fmt.Print(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)

	greeting = "Hello,"
	name = "Amanda"
	year = 2022

	temp := fmt.Sprintln(greeting, name, year)

	fmt.Println(temp)

	s := fmt.Sprintf("%#x\t%b\t%x", y, y, y)

	fmt.Println(s)
}
