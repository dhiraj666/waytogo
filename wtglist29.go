package main

import (
	"fmt"
)

func main() {
	s := "good bye"
	var p *string

	p = &s
	*p = "hello "
	fmt.Printf("Here is the pointer p: %p\n", p)
	fmt.Printf("Here is the string *p: %s\n", *p)
	fmt.Printf("Here is the sring s : %s\n", s)
}
