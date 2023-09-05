package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0

	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}

func f2() (r int) {
	r = 1
	return
}

func f3() (int, int) {
	return 5, 6
}

func variadicFunctions() {
	
}

func main() {
	// xs := []float64{98, 93, 77, 82, 83}
	// fmt.Println(average(xs))
	a := f2()
	fmt.Println(a)
}