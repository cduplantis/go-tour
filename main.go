package main

import (
	"fmt"
)

const (
	Big   = 1 << 100 // high precision number
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 - 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Printf("needInt(Small): %v\n", needInt(Small))
	fmt.Printf("needFloat(Small): %v\n", needFloat(Small))
	fmt.Printf("needFloat(Big): %v\n", needFloat(Big))
}
