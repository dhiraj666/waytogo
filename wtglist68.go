/*
 func typecheck(..,..,values ...interface{}) {
	for _, value := range values {
	switch v := value.(type){
	case int :
	case float :
	case string:
	case bool:
	default
	}
	}
 }

*/

// defer and tracing

// defer keyword allow us to postpone the execution of a statement or a function until the end of the enclosing function
// it executes something  when the enclosing function returns

package main

import (
	"fmt"
)

func main() {
	Function1()
	a()
	f()
}

func Function1() {
	fmt.Printf("In Function1 at the top\n")
	defer Function2()
	fmt.Printf("In Function1 at the bottom\n")
}

func Function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!")
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// defers are executed like a stack or LastInFirsOut

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}

// defer allow us to guarantee that certain clean up tasks are performed before we return from a function

// closing a file stream defer file.Close()
// unlocking a locked resource defer mu.Unlock()
// Printing a footer in a report defer printFooter
// closing a database connection defer disconnectFromDB()
