package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch)
	go Walk(t2, ch2)
	match := true
	for i := 0; match && i < 10; i++ {
		v, ok := <-ch
		v2, ok2 := <-ch2
		match = ok && ok2 && v == v2
	}
	return match
}

func main() {
	// fmt.Printf("tree.New(1).String(): %v\n", tree.New(1).String())
	// ch := make(chan int)
	// go Walk(tree.New(1), ch)
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%d ch: %v\n", i, <-ch)
	// }
	test1 := Same(tree.New(1), tree.New(1))
	fmt.Printf("test1: %v\n", test1)
	test2 := Same(tree.New(1), tree.New(2))
	fmt.Printf("test2: %v\n", test2)
}
