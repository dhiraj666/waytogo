// for loop

package main

import (
	"fmt"
)

func main() {
	var G string = "G"
	for r := 0; r < 25; r++ {
		for c := 0; c <= r; c++ {
			fmt.Printf("%s", G)
		}
		fmt.Printf("\n")
	}
}
