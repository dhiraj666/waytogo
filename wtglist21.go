// presuffix.go

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "The quick brown fox jumps over a lazy dog"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	fmt.Printf("String contains %t", strings.Contains(str, "jumps"))
}
