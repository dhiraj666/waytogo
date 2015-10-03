// package also has a strings.NewReader function

// Read() to read a []byte
// ReadByte() and ReadRune() to read the next byte or rune from the string

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "666"

	var an int
	var news string

	fmt.Printf("The size of ints is %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Printf("The Integer is : %d\n", an)

	an = an + 5
	news = strconv.Itoa(an)
	fmt.Printf("The new string is %s\n", news)

}
