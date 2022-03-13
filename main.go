package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return  -1
}

func main() {
	si := []int { 10, 20, 15, -10, 0 }
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "baz", "qux"}
	fmt.Println(Index(ss, "hello"))
}
