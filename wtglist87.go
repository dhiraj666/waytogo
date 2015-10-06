//  0 1 1 2 3 5
package main

import (
	"fmt"
)

var (
	a, b, c int
)

func main() {
	fibonacciImperative(15)
}

func fibonacciImperative(n int) {
	a = 0
	b = 1
	c = a + b
	fmt.Printf("%d,%d,", a, b)
	for i := 2; i < n; i++ {
		fmt.Printf("%d,", c)
		a = b
		b = c
		c = a + b
	}

}
