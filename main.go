package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// send sum to c
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s)
	// make a channel
	c := make(chan int)
	// split the array in two parts
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// receive from c
	x, y := <-c, <-c
	// print the sum
	fmt.Println(x, y, x+y)
}
