//  passing a variable number of parameters
/*
If the last parameter of a function is followed by ... type , this indicates that the function can deal
with a variable number of parameters of that type , possibly also 0
called variadic functions

*/

package main

import (
	"fmt"
)

func main() {
	x := Min(1, 3, 2, 0)
	fmt.Printf("The Minimum is : %d\n", x)

	arr := []int{7, 9, 3, 5, 1}
	//  check how array is passed to the Min function which accepts variadic parameters
	x = Min(arr...)
	fmt.Printf("The Minimum in the array is  : %d\n", x)
}

func Min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
