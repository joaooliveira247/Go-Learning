package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	// create a pointer w/ type given
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}