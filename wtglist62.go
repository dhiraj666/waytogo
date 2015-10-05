// error return value

package main

import "fmt"

func main() {
	s, err := MySqrt(6.0)
	if err == nil {
		fmt.Println(s)
	}
	s, err = MySqrt(-3.0)
	if err == nil {
		fmt.Println(s)
	}
}

func MySqrt(f float64) (s float64, err error) {
	if f > 0 {
		s = f * f
		return s, nil
	} else {
		return 0, err
	}

}
