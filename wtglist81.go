/*
	A closure is a function that captures some external state
	for example , the state of the function in which it is created
	Closure inherits the scope of the function in which it is declared

	Closure can be used in performing clean error checking
*/
package main

// function returning another function
import (
	"fmt"
)

func main() {
	p2 := Add2()
	fmt.Printf("%v\n", p2(3))

	TwoAdder := Adder(2)
	fmt.Printf("%v\n", TwoAdder(3))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return b + a
	}
}
