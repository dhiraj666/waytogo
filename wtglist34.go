// error is nil , otherwise it contains the error information

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orgn string = "ABC"

	var an int

	var err error
	an, err = strconv.Atoi(orgn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The integer is %d\n", an)

}
