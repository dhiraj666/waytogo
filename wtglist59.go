// When it becomes necessary to return more than
// 4 or 5 values from a func
// it is better to use slice if the values are of same type ( homogenous)
// or pointer to a struct type if it is of different type

package main

import (
	"fmt"
)

func main() {
	var a, b int = 5, 2
	DoSomething(&a)
	DoSomethingDiff(b)
}

func DoSomething(a *int) {
	b := a
	fmt.Println("dosomethingpointer\n", b)

}

func DoSomethingDiff(a int) {
	b := &a
	fmt.Println("dosomethingvalue\n", b)
}
