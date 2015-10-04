package main

import (
	"fmt"
)

// switch 2nd form different way

func main() {

	var num1 int = 30
	switch {
	case num1 < 0:
		fmt.Println("Number is negative")
	case num1 > 0:
		fmt.Println("Number is positive")
	default:
		fmt.Println("don't know !!!!")
	}
}
