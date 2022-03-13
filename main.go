package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Printf("%v\n", Vertex{1, 2})
}
