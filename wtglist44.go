// condition controlled
// break only exits from the loop
// return exits from the function in which loop is coded
package main

import (
	"fmt"
)

func main() {
	var i int = 5
	for i >= 0 {
		i = i - 1
		fmt.Printf("The variable i is now %d\n", i)
	}
}

// infinite loop can be used in server function where it is waited for incoming requests
