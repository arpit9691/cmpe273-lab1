package Perimeter

import "testing"
import "fmt"

func Test_PerimeterCircle(t *testing.T) {

	var s Shaper
	c := Circle{r: 4}
	s = c

	if s.Perimeter() < 26 {

		t.Log("Fail")
	} else {

		t.Log("Pass")
		fmt.Println("1")
	}
	fmt.Println("Perimeter of Shaper(Circle): ", s.Perimeter())
	fmt.Println()
}

func Test_PerimeterRectangle(t1 *testing.T) {

	var s Shaper
	r := Rectangle{l: 6, w: 3}
	s = r

	if s.Perimeter() != 18 {

		t1.Log("Fail")
	} else {

		t1.Log("Pass")
		fmt.Println("2")
	}
	fmt.Println("Perimeter of Shaper(Rectangle): ", s.Perimeter())
	fmt.Println()
}
