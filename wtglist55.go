// function calling in another function call

package main

import (
	"fmt"
)

func main() {
	// calling functionone with another function as argument
	functionOne(functiontwo(3, 4))

}

func functionOne(a, b, c int) {
	//return 5, 6, 7
	fmt.Println("lolgllakdjf")
}

func functiontwo(d, e int) (int, int, int) {
	return 8, 9, 10
}
