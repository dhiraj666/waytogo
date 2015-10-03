// package time gives us a datatype
// time.Time (to be used as value)
// functionality for displaying and measuring time and dates

// current time is given by time.Now() and the parts of a time can then be obtained as
//t.Day(),
//t.Minute(),
//t.Month()
//t.Year()
// type Duration represents the elapsed time between two instants as an int64 nanosecond count

// time.go

package main

import (
	"fmt"
	"time"
)

// one can declare a global variable and not used it
var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)

	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())

	t = time.Now().UTC()
	fmt.Println(t)

	week = 60 * 60 * 24 * 7 * 1e9 //must be in nanosec

	week_from_now := t.Add(week)

	fmt.Println(week_from_now)

	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	s := t.Format("20060102")

	fmt.Println(t, "=>", s)

}
