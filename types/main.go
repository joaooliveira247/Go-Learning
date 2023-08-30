package main

import "fmt"

func integers() {
	// int type
	//(uint8 = byte, uint16, uint32, uint64) - Unsigned int [0, +bit]
	//(int8, int16, int32 = rune, int64) - int [-bit, +bit]
	fmt.Println("1 + 1 =", 1 + 1)
}

func float() {
	/*
	float type
	float32
	float64 - double precision
	*/
	fmt.Println("1.0 + 1.0 =", 1.0 + 1.0)
}

func operators() {
	/*
	+ -> adition 
	- -> subtracion
	* -> multiplication
	/ -> division
	% -> rest
	another functions import "math" package
	*/
	fmt.Println("rest of 3 / 2 =", 3 % 2)
}

func strings() {
	fmt.Println(len("Hello, World!"))
	fmt.Println("Hello, World!"[1]) // slicing
	fmt.Println("Hello" + "World!") // concant	
}

func booleans() {
	/*
	Booleans Operatos
	&& - And
	|| - Or
	! - different
	*/
	println(true && true)
	println(true && false)
	println(true || true)
	println(true || false)
	println(!false)
}

func main() {
	booleans()
}