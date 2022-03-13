package main

import (
	"fmt"
)

const Pi = 3.14

func main() {
	const world = "Casey"  // don't use := for const assignments
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day") // Println concatenates with spaces

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
