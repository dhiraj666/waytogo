package main

import (
	"fmt"
)

func main() {
	g := func() {
		println("Helloworld")
	}
	g()
	fmt.Printf("%T", g)
}
