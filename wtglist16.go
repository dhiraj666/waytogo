// there is no operator overloading
// If they are interfaces
// they must implement the same interface type

// %t is used as a format specifier for booleans

// int type exists but float type doesn't exist

// float32 is reliably accurate to about 7 decimal places
// float64 is reliabley accurate to 15 decimal places

// all the functions of math type expects float64

// mixing of constants are ok

//safedownsizing

package main

import (
	"fmt"
	"math"
)

func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range ", n)
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)

		if fraction >= 0.5 {
			whole++
		}

		return int(whole)

	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}
