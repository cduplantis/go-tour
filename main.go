package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Print("OS X")
	case "linux":
		fmt.Print("Linux.")
	default:
		fmt.Printf("%s!\n", os)
	}
}
