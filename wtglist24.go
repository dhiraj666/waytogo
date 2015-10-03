package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "The quick brown fox jumps over the lazy dog"
	s1 := strings.Fields(str)

	fmt.Printf("Splitted in slice : %v\n", s1)

	for _, val := range s1 {
		fmt.Printf("%s - ", val)
	}

	fmt.Println()

}
