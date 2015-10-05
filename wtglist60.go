// named return variables

/*
	Multiple return.go
	When there is more than 1 unnamed return variable ,
	they must be enclosed within ()

	Named variables used as result paramters are automatically
	initialized to their zero value
	a simple return statement is sufficient
	furthermore even when there is only 1 named return
	variable , it has to be put inside ()


*/

package main

import (
	"fmt"
)

var num int = 10
var numx2, numx3 int

func main() {
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input

	return
}
