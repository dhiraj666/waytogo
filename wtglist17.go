package main

import (
	"fmt"
)

var a complex64 = 5 + 10i

func main() {
	fmt.Printf("The value is %v", a)

	d := complex(5, 12)
	fmt.Printf("%v", real(d))
	fmt.Printf("%v", imag(d))
}
