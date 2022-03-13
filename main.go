package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

const epsilon = 1e-8 // or any small value you see fit

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	// loop to calc the square root of x
	z := 1.0
	p := z + epsilon
	for i := 0; i < 10 ; i++ {
		z -= (z*z - x) / (2 * z)
		if almostEqual(p, z) {
			break
		}
		p = z
	}
	return z, nil
}

func main() {
	report(1)
	report(2)
	report(3)
	report(4)
	report(14)
	report(16)
	report(225)
	report(-81)
}

func report(i float64) {
	s, e := Sqrt(i)
	s2 := math.Sqrt(i)
	fmt.Printf("Sqrt(%v) = %v, math.Sqrt(%v) = %v, %v\n", i, s, i, s2, e)
}
