package main

import (
	"fmt"
)

type Stack [10]int

func (st *Stack) Pop() int {
	v := 0
	for i := len(st) - 1; i >= 0; i-- {
		if v = st[i]; v != 0 {
			st[i] = 0
			return v
		}

	}
	return 0
}

// To declare a function implemented outside Go , such as an assembly routine
// You simply give the name and signature , and no body

/*
	func flushICache(begin, end uintptr) //implemented externally

*/

// type binOp func(int, int ) int in this case also binOp is omitted

// Functions are first class values  , they can be assigned to a variable like in add := binOp

// Go has now no concept of generics

// defining one function which can be applied to a number of variable types

// However most cases can be solved simply by using interfaces , especially the empty interface and a type switch

// or by using reflection

// Code complexity is increased using empty interface or type switch or reflection , for better performance and more readable code , create the function explicitly for every type used

func main() {
	var st Stack
	fmt.Println(st.Pop())
}
