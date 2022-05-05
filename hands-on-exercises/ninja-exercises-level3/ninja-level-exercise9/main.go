package main

import "fmt"

func main() {
	favSport := "Swimming"

	switch favSport {
	case "Soccer":
		fmt.Println("This is not my fav sport")
	case "Volleyball":
		fmt.Println("This is also not my fav sport")
	case "Swimming":
		fmt.Println("This is my fav sport")
	}
}
