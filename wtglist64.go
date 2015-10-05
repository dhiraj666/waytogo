// min max between two numbers

package main

import (
	"fmt"
)

func main() {
	var min, max int
	min, max = MinMax(89, 98)
	fmt.Printf("Minimum is %d : AND Maximum is %d: ", min, max)

}

func MinMax(a, b int) (min int, max int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return
}
