package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}


func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3,4}

	a = f
	a = &v

	// this line doesn't compile because the receiver in Abs() is a pointer
	// either change the receiver to accept a value or pass a pointer
	// a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y int
}

func (v *Vertex) Abs() float64 {
	// the commented line, 19, doesn't compile because this accepts a pointer
	// either change the receiver to accept a value or pass a pointer
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

