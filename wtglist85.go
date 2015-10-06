package main

import (
	"fmt"
)

func main() {
	a := [...]string{"a", "b", "c", "d", "e"}
	for i := range a {
		fmt.Println("array item ", i, "is", a[i])
	}
}
