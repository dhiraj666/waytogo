// func accepts 2 integers
// returns sum product and diff
// named return aswell as unnamed return

package main

import (
	"fmt"
)

var sum, prod, diff int

func main() {
	sum, prod, diff = UNamedReturn(6, 4)
	PrintValues()
	sum, prod, diff = NamedReturn(6, 4)
	PrintValues()
}
func PrintValues() {
	fmt.Printf("sum =%02d\n", sum)
	fmt.Printf("Product = %02d\n", prod)
	fmt.Printf("Difference = %02d\n", diff)
}
func UNamedReturn(a, b int) (int, int, int) {

	return a + b, a * b, b - a
}

func NamedReturn(a, b int) (sum int, prod int, diff int) {
	sum = a + b
	prod = a * b
	diff = b - a

	return
}
