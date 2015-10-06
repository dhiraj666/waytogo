/*
	When doing heavy computation, we should not repeat any calculation that has already been done
	Instead cache the calculated value in memory called memoization
	or greater avalanche of recalculations

	// simple fibonacci took 4.73 seconds
	but optimized fibonacci took .001000s

*/
// For me , it took 3 minute 23 sec here for 1e6 values of fibonacci

// fibonacci memoization

package main

import (
	"fmt"
	"time"
)

const LIM = 1e6

var fibs [LIM]uint64

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonaccci(i)
		//	fmt.Printf("fibonaccci(%d) is %d\n", i, result)
	}

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("Long calculation took this amount of time : %s\n", delta)
}

func fibonaccci(n int) (res uint64) {
	// memoization : check if fibonacci(n) is already known in array
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonaccci(n-1) + fibonaccci(n-2)
	}
	fibs[n] = res
	return
}
