// index in string

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hi , I am Dhiraj , Hi."
	fmt.Printf("The position of \"Dhiraj\" is :")
	fmt.Printf("%d\n", strings.Index(str, "Dhiraj"))
	fmt.Printf("The position of the first instance of \"Hi\" is : ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is : ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	// replacing a string strings.Replace(str,old,new,n) and if n= -1 , all occurrences are replaced

	var orgstring = "Hello, Dhiraj Das Dhiraj Das Dhiraj Das Dhiraj Das Dhiraj Das"
	var manyG = "ggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg"
	strings.Replace(orgstring, "h", "hhh", -1)
	fmt.Printf("Number of H's in %s is : ", orgstring)
	fmt.Printf("%d\n", strings.Count(orgstring, "h"))
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))
}
