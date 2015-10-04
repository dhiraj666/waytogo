// one cannot take the address
// of the literal or a constant
// go is memory safe and pointer arithmetic is not possible
package main

import (
	"fmt"
)

const (
	i = 3
)

var ptr *int

func main() {
	ptr = &i
	fmt.Printf("%p", ptr)
}

// pointers life time is independent of the scope in which they are created
