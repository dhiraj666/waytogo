// Functions can be used as parameters
// callback function

package main

import (
	"fmt"
)

func main() {
	callback(1, Add)
}

func Add(a, b int) {
	fmt.Printf("The sum is %d and %d is : %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) //  this becomes add(1,2)
}
