package main

import (
	"fmt"
)

func main() {
	arraypt(&[3]int{1, 1 * 1 * 1, 1 * 2 * 3})
}

func arraypt(arrayv *[3]int) {
	fmt.Println(arrayv)
}
