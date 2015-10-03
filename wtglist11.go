//Memory in a computer is used in program as a enormous number of boxes called words
//All words have the same length 32 bits or 64 bits
//All words are identified by their memory address

// primitive types are value types , they point directly to their value contained in the memory

// Variables of value type are contained in stack memory

package main

import (
	"fmt"
)

var (
	a int = 2
	b int = a
)

func main() {
	fmt.Println(a, b)
	fmt.Println(&a, &b)
}
