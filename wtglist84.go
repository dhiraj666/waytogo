package main

/*
	Data Structures :
	Collections
	array
	slices
	maps

*/
import "fmt"

// Array is homogenous numbered
// can be primitive aswell as self defined types
// length must be constant
// it is also the type
// so [5]int [10]int are different

// INitialization of an array with values known at compile time is done with array line
// any type is empty interface type

// maximum array length is 2 Gb

// declaration format
// var identifier [len]type

// indexoutofbound is either runtime or compile time

func main() {
	var arrayvar [5]int
	for i := 0; i < len(arrayvar); i++ {
		arrayvar[i] = i * 2
	}
	for i := 0; i < len(arrayvar); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arrayvar[i])
	}
}
