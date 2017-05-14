package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e)) 
	}
	return ""
}

func Sqrt(x float64) (float64, error) {
	err := ErrNegativeSqrt(x).Error()
	if err != "" {
		return x, ErrNegativeSqrt(x)
	}
	z := 1.0
	for v := 0; v < int(x); v++ {
		z -= (math.Pow(z, 2) - x)/(2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
