/*
	A primitive but effective ways of
	tracing the execution of a program
	is printing a message when entering and leaving certain functions
	This can be done with the following two functions
	func trace(s string) {fmt.Println("entering", s)}
	func untrace(s string){fmt.Println("leaving ", s)}
*/

package main

import (
	"fmt"
)

func trace(s string) {
	fmt.Println("entering:", s)
}

func untrace(s string) {
	fmt.Println("leaving:", s)
}

func ab() {
	trace("ab")
	defer untrace("ab")
	fmt.Println("in ab ")
}

func cd() {
	trace("cd")
	defer untrace("cd")
	fmt.Println("in cd")
	ab()
}

func main() {
	cd()
}
