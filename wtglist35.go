package main

import (
	"fmt"
)

func main() {
	var switchvariable int

	switch switchvariable {
	case 0:
		fmt.Println("The value is 0")
	case 1:
		fmt.Println("This case will not be executed ")
	default:
		fmt.Println("value")
	}
}
