// warning the code in line 17 will cause the stack to overflow
// so beware , it is just for illustrative purposes

package main

import (
	"fmt"
)

func main() {
	fmt.Println(printrec(10))
}

func printrec(n int) (r int) {
	if n <= 1 {
		r = n
	} else {
		r = printrec(n-1) + printrec(n)
	}
	return r
}
