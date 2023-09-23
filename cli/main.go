package main

import (
	"fmt"
	"flag"
	"math/rand"
)

func main() {
	maxp := flag.Int("max", 6, "the max value")
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(*maxp)
	fmt.Println(rand.Intn(*maxp))
}