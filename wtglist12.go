//reference types in golang are
// pointers slices maps channels
// variables  that are referenced are stored in the heap which is garbage collected
// which is a much larger memory space than the stack

package main

import (
	"fmt"
	"math"
)

var Pi float64

func init() {
	Pi = 4 * math.Atan(1) // init() function computes Pi
}

func swap() {
	var a, b = 2, 3
	fmt.Println("Before swapping", a, b)
	a, b = b, a

	fmt.Println("after swapping", a, b)
}

func main() {
	swap()
	fmt.Println(Pi)
}

/*
	func init() {
		go backend()
	}

*/
