// passing a big array to a function quickly uses up much memory

// There are two solutions to prevent this
/*
	1) PASS A POINTER TO THE ARRAY
	2) USE A SLICE OF THE ARRAY


*/

package main

import (
	"fmt"
)

func main() {
	array := [3]float64{1.2, 3.4, 2.1}
	x := sum(&array)
	fmt.Printf("The sum of the array is %.2f", x)
}

func sum(array *[3]float64) (x float64) {
	//sumvar := 0.0
	var summ float64
	for i := 0; i < len(array); i++ {
		summ += array[i]
	}
	return summ
}
