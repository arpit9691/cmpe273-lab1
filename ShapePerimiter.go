package main

import (
	"fmt"
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

func main() {
	r := Rectangle{l: 6, w: 3}
	fmt.Println("Rectangle details are: ", r)

	var s Shaper
	s = r 
	fmt.Println("Area of Shaper(Rectangle): ", s.Area())
	fmt.Println()

	fmt.Println("Perimeter of Shaper(Rectangle): ", s.Perimeter())
	fmt.Println()

	c := Circle{r: 4}
	fmt.Println("Circle details are: ", c)
	s = c 
	fmt.Println("Area of Shaper(Circle): ", s.Area())
	fmt.Println()

	fmt.Println("Perimeter of Shaper(Circle): ", s.Perimeter())
	fmt.Println()
}
