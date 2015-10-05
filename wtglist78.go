package main

import (
	"fmt"
)

func main() {
	func() {
		sum := 0
		for i := 0; i < 1e6; i++ {
			sum += i
		}
		fmt.Println(sum)
	}()

}
