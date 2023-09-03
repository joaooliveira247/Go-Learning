package main

import "fmt"

func array() {
	var x [5]int

	x[4] = 100

	fmt.Println(x)
}

func anotherArrayExample() {
	var x [5]float64
	
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0

	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Println(total / float64(len(x)))
}

func arrayWithRange() {
	x := [5]float64{
		98, 93, 77, 82, 83,
	}

	var total float64 = 0

	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))
}

func emptySlices() {
	var x []float64
	fmt.Println(x)
	y := make([]float64, 5)
	fmt.Println(y)
	z := make([]float64, 5, 10)
	fmt.Println(z)
}

func normalSlices() {
	x := []float64{1, 2, 3, 4, 5}
	fmt.Println(x[:3])
}

func sliceAppend() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := append(slice1, 4, 5, 6)

	fmt.Println(slice1, slice2)
}

func sliceCopy() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func emptyMap() {

	// This map isn't load
	var x map[string]int

	fmt.Println(x)
}

func loadedMap() {
	x := make(map[string]int)

	x["key"] = 10

	fmt.Println(x, x["key"])
}

func elementsExample() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	delete(elements, "Ne")

	fmt.Println(elements)

	name, ok := elements["Un"]

	fmt.Println(name, ok)
}

func main() {
	elementsExample()
}