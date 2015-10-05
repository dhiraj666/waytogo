// closures

package main

/*
	function literals

*/
import (
	"fmt"
)

// anonymous func , lambda func , func literal, closure

// closures can be assigned to a variable which is a reference to that function

/*

	fplues := func(x, y int ) int {return x+y }
*/

func main() {
	// one way to invoke is to do it directly
	// and another way is to first declare and then use the variable as a function call

	z := func(x, y int) int {
		return x + y
	}

	//fmt.Println(z)
	fmt.Println(z(8, 12))

	//fmt.Println(z)
}
