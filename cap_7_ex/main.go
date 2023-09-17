package main

import (
	"fmt"
	"math"
)

type Circle struct {
	a, b, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func main() {
	c := Circle{0, 0, 5}
	res := fmt.Sprintf("Area: %.2f | Perimeter %.2f"	, c.area(), c.perimeter())
	fmt.Println(res)
}
