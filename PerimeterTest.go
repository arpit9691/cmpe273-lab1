package Perimeter

import (
	//"fmt"
	"math"
)

type Shaper interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	l, w float64
}

func (r Rectangle) Area() float64 {
	return r.l * r.w
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.l + r.w)
}

type Circle struct {
	r float64
}


func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}
