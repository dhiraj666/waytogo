package main

import (
	"fmt"
)

func main() {
	s := ""
	for s != "aaaaa" {
		fmt.Println("value of s is :", s)
		s = s + "a"
	}
}
