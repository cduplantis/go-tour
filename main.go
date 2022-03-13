package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)
	for _, w := range words {
		if v, ok := counts[w]; ok {
			counts[w] = v + 1
		} else {
			counts[w] = 1
		}
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
