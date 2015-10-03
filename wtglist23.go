package main

import (
	"fmt"
	"strings"
)

func main() {
	var stringnew = "dhiraj"
	fmt.Printf("%s\n", strings.Repeat(stringnew, 4))
	fmt.Printf("%s\n", strings.ToUpper(stringnew))

	// Trimming space

	var trimSpace = "		      dhiraj is learning golang and dhiraj you have to love yourself			"
	fmt.Printf("%s", strings.TrimSpace(trimSpace))

	// strings.Trim(str, "\r\n")

}
