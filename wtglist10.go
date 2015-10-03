// operating system on which it runs

package main

import (
	"fmt"
	"os"
)

func main() {
	var goos string = os.Getenv("GOPATH")
	fmt.Println("The GOPATH is %s", goos)
}
