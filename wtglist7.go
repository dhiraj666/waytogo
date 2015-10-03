// go template

package main

import (
	"fmt"
)

const (
	// can be boolean
	a bool = true
	//can be integer float or complex

	in int     = 3
	f  float64 = 3.14
	// can be string
	dhiraj string = "dhiraj kumar das "
)

const c = "carat"

var v int = 5

type T struct{}

func init() { // initialization of package

}

func main() {

	Func1()
	fmt.Println(a, 7in, f, dhiraj)
}
func (t T) Method1() {

}

func Func1() {

}
