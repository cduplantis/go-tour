package main

import (
	"fmt"
)

type List[T any] struct {
	next *List[T]
	val  T
}

func NewList[T any](val T) *List[T] {
	return &List[T]{val: val}
}

// implement a single linked list
func (l *List[T]) Add(val T) {
	if l.next == nil {
		l.next = &List[T]{val: val}
	} else {
		l.next.Add(val)
	}
}

func (l *List[T]) Get(i int) T {
	if i == 0 {
		return l.val
	}
	if l.next == nil {
		panic("index out of range")
	}
	return l.next.Get(i - 1)
}

func (l *List[T]) Len() int {
	if l.next == nil {
		return 0
	}
	return 1 + l.next.Len()
}

func (l *List[T]) Print() {
	if l == nil {
		fmt.Println()
		return
	}
	fmt.Print(l.val, " ")
	l.next.Print()
}

func (l List[T]) String() string {
	if l.next == nil {
		return fmt.Sprintf("%v", l.val)
	}
	return fmt.Sprintf("%v -> %v", l.val, l.next)
}

func main() {
	l := NewList(0)
	for i := 2; i < 10; i += 2 {
		l.Add(i)
	}
	for i := 1; i < 10; i += 2 {
		l.Add(i)
	}
	fmt.Printf("%s\n", l)
	l.Print()
	fmt.Printf("len: %d\n", l.Len())

	fmt.Printf("l[0]: %d\n", l.Get(0))
	fmt.Printf("l[0]: %d\n", l.Get(4))
}
