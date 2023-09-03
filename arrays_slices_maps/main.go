package main

import (
	"fmt"
	"sort"
)

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

func anotherMapAssignment() {
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	fmt.Println(elements)
}

func mapsInMaps() {
	elements := map[string]map[string]string{
		"H":  map[string]string{"name": "Hydrogen", "state": "gas"},
		"He": map[string]string{"name": "Helium", "stage": "gas"},
		"Li": map[string]string{"name": "Lithium", "stage": "solid"},
		"Be": map[string]string{"name": "Beryllium", "stage": "solid"},
		"B":  map[string]string{"name": "Boron", "stage": "solid"},
		"C":  map[string]string{"name": "Carbon", "stage": "solid"},
		"N":  map[string]string{"name": "Nitrogen", "stage": "gas"},
		"O":  map[string]string{"name": "Oxigen", "stage": "gas"},
		"F":  map[string]string{"name": "Fluorine", "stage": "gas"},
		"Ne": map[string]string{"name": "Neon", "stage": "gas"},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el, ok)
	}
}

func ex1() {
	xArray := [5]int8{
		1, 2, 3, 4, 5,
	}
	fmt.Println(xArray[3])

	xSlice := []int8{
		1, 2, 3, 4, 5,
	}
	fmt.Println(xSlice[3])

}

func ex2() {
	xSlice := make([]int, 3, 9)
	fmt.Println(xSlice)
	fmt.Println(len(xSlice))
}

func ex3() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])
}

func ex4() {
	x := []int{
		48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,
	}

	lowest := x[0]

	// using for
	for _, value := range x {
		if value < lowest {
			lowest = value
		}
	}
	fmt.Println("for | ", lowest)

	// using sort
	sort.Ints(x)
	fmt.Println(x[0])
}

func main() {
	ex4()
}
