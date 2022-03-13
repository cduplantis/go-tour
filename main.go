package main

import (
	"fmt"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	pow := make([]int, 10) // [0 0 0 0 0 0 0 0 0 0]

	for i := range pow {
		pow[i] = 1 << uint(i) 
	}

	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}
