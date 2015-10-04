// for range construct

// strings is a collection of unicode characters or runes so it can be applied for strings too

package main

import (
	"fmt"
)

func main() {
	str := "Go is a beautiful language"

	for pos, char := range str {
		fmt.Printf("Character on position %d is : %c\n", pos, char)
	}
	fmt.Println()

	str2 := "日本語"
	for pos, char := range str2 {
		fmt.Printf("Character %c starts at byte postion %d\n", char, pos)
	}
	fmt.Println()

}
