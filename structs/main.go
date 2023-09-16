package main

import (
	"fmt"
	"math"
)

type Circle struct {
	a, b, r float64
}

func emptyStruct() {
	c := new(Circle)
	fmt.Println(c)
}

func newPointerStruct() {
	c := &Circle{0, 0, 5}
	fmt.Println(c)
}

func fieldAcessStruct() {
	c := &Circle{0, 0, 5}
	fmt.Println(c)
	c.a = 10
	c.b = 16
	fmt.Println(c.a, c.b, c.r)
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x2
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}


// methods

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}


type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w 
}

func main() {
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}