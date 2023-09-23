package main

import (
	"create_packages/math"
	"fmt"
)

func main() {
	a := []float64{
		1, 2, 3, 4, 5, 6,
	}

	fmt.Println(math.Avg(a))
	fmt.Println(math.Square(13))
}