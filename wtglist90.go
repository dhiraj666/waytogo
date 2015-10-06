// Multidimensional arrays

// Arrays are always one dimensional but they may be composed to form multidimensional arrays

/*
	[3][4]int

*/

package main

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func main() {
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			screen[i][j] = 0
		}
	}
}
