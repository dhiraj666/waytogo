/*
	Make a Function that has arguments a variable number of ints and which prints
	each integer on a separate line
*/
package main

import (
	"fmt"
)

func varargs(a ...int) {
	for i, _ := range a {
		fmt.Printf("%d\n", a[i])
	}
}

func main() {
	arr := []int{2, 3, 4, 5}
	varargs(arr...)
}
