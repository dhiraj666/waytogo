// recursive functions
// fibonacci series
// recursive solution can be used in quicksort algorithm also
package main

import (
	"fmt"
)

func main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is %d\n", i, result)
	}
}

func fibonacci(f int) (res int) {
	if f <= 1 {
		res = 1
	} else {
		res = fibonacci(f-1) + fibonacci(f-2)
	}
	return
}
