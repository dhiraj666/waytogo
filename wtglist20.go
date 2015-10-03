// characters are type of integers

// characters are unicode points or runes
// unicode characters is represented by an int in memory
package main

import (
	"fmt"
)

var (
	ch  = 'a'
	cha = 65
)

func main() {
	fmt.Println(ch)
	fmt.Printf("%c", cha)

}
