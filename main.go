package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 5, 7, 11, 13}

	s = s[1:4] // 2,3,5
	fmt.Println(s)

	s = s[:2] // 2,3
	fmt.Println(s)

	s = s[1:] // 3
	fmt.Println(s)

}
