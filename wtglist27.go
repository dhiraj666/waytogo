// programmer controls over the basic memory layout
// pointers are important for performance and indispensable

// program store values in memory
// each memory block ( or word ) has an address
// represented as hexa decimal number

// &operator which when placed before a variable give us the memory address
// of that variable

package main

import (
	"fmt"
)

var i1 = 5

var a *int

func main() {
	fmt.Printf("An integer %d, it's location in memory is %p\n", i1, &i1)
	fmt.Println(a)
}
