package main

import (
	"fmt"
)

var i = 5

func main() {
	for {
		i = i - 1
		fmt.Printf("The variable is now %d\n", i)
		if i < 0 {
			break

		}

	}
}
