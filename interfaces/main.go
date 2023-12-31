package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	area() float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x2
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func totalArea(shapes ...Shape) float64 {
	total := 0.0
	for _, v := range shapes {
		total += v.area()
	}
	return total
}

// interface like fields

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	total := 0.0
	for _, v := range m.shapes {
		total += v.area()
	}
	return total
}

func main() {
	// c1 := Circle{0, 0, 5}
	// c2 := Circle{0, 0, 6}
	// r1 := Rectangle{0, 0, 10, 10,}
	// fmt.Println(totalArea(&c1, &c2, &r1))
	multiShape := MultiShape{
		shapes: []Shape{
			&Circle{0, 0, 5},
			&Rectangle{0, 0, 10, 10},
		},
	}
	fmt.Println(multiShape.area())
}
