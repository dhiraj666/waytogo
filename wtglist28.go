// pointer.go

package main

import (
	"fmt"
)

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

	var intp *int
	intp = &i1

	fmt.Printf("The value at memory location %p is %d\n", intp, *intp)

}
