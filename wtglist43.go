// fizz buzz problem

package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 30; i++ {
		fmt.Printf("%d\n", i)
		if i%3 == 0 {
			fmt.Println("fizz")
		}
		if i%5 == 0 {
			fmt.Println("buzz")
		}
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		}
	}
}
