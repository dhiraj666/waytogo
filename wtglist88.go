// Array Literals

package main

import "fmt"

func main() {
	var arrAge = [5]int{24, 25, 31, 50, 64}
	var arrLazy = [...]int{5, 6, 7, 8, 22}
	var arrkeyValue = [5]string{3: "dhiraj", 4: "das"}

	for i := 0; i < len(arrkeyValue); i++ {
		fmt.Printf("Person at %d is %s \n", i, arrkeyValue[i])
	}
	for i := 0; i < len(arrAge); i++ {
		fmt.Printf("Person at %d is %d \n", i, arrAge[i])
	}

	for i := 0; i < len(arrLazy); i++ {
		fmt.Printf("Person at %d is %d \n", i, arrLazy[i])
	}

}
