package main

import (
	"fmt"
	"strings"
)

func mekeEvenGenerator() func () uint  {
	i := uint(0)
	return func() (ret uint)  {
		ret = i
		i += 2
		return
	}
}

func main() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 1))
	fmt.Println(strings.Repeat("-", 20))

	x := 0

	increment := func() int {
		x++
		return x
	}
	fmt.Println("X value is:", x)
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println("X value now is:", x)

	fmt.Println(strings.Repeat("-", 20))

	nextEven := mekeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

}
