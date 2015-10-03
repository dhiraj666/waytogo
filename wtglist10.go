// operating system on which it runs

package main

import (
	"fmt"
	"os"
)

func main() {
	var goos string = os.Getenv("GOOS")
	fmt.Println("The operating system is %s", goos)
}
