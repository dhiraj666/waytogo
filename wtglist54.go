// Functions are the basic building blocks in gocode
// DRY - Don't Repeat Yourself
// } or return statement
// return can be used to end an infinite for loop or stop a gorutine

/*
	There are three types of functions in Go
	Normal functions with an identifier
	Anonymous or lambda functions
	methods
*/

// any of these can have parameters and return values

package main

func main() {
	println("In main before calling greeting")
	greeting()
	println("In main after calling greeting")
}

func greeting() {
	println("In greeting!!")
}
