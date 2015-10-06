// pointer array
// when assinging an array to another
// a distinct copy in memory of the array is made

package main

import (
	"fmt"
)

func main() {
	var array [3]int
	var secondarray [3]int
	fpppointer(&array)
	fppvalue(array)
	secondarray = array
	fmt.Printf("%p\n", &secondarray)
	fmt.Printf("%p\n", &array)
}

func fppvalue(array [3]int) {
	fmt.Println(array)
}

func fpppointer(array *[3]int) {
	fmt.Println(array)
}
