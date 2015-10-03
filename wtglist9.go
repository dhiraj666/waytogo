// declaration of variable
// var identifier type = value

package main

import (
	"fmt"
	"os"
)

var (
	a      int    = 15
	b      bool   = false
	c      string = "Go says hello to the world !!"
	d             = "dhiraj das"
	HOME          = os.Getenv("HOME")
	USER          = os.Getenv("USER")
	GOROOT        = os.Getenv("GOROOT")
	GOPATH        = os.Getenv("GOPATH")
	GOARCH        = os.Getenv("GOARCH")
)

func main() {
	fmt.Println(a, b, c, d)
	fmt.Println("Home, USER, GOROOT, GOPATH, GOARCH", HOME, USER, GOROOT, GOPATH, GOARCH)
}
