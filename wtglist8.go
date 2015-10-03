// constant using iota

package main

import "fmt"

type Weekdays int

const (
	Sunday Weekdays = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	var a Weekdays
	for a < Saturday {
		fmt.Println(a)
		if a == Saturday {
			break
		}
		a++
	}

}
