// reference types like
// slices , maps , interfaces and channels are pass by reference by default

// printing to the console , sending a mail , logging an error

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", Multiply3Nums(2, 5, 6))

}

func Multiply3Nums(a, b, c int) int {
	return a * b * c
}
