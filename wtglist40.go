// string for

package main

import (
	"fmt"
)

func main() {
	str := "Go is a beautiful programming language!!"
	fmt.Printf("The length of str is %d\n", len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("Character on position %d is : %c\n ", i, str[i])
	}
}
