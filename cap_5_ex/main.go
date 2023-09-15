package main

import (
	"fmt"
)

func sum(slice []int) int {
	total := 0
	for _, v := range slice {
		total += v
	}
	return total
}

func half(x int) (int, bool) {
	if x%2 == 0 {
		return x / 2, true
	}
	return x / 2, false
}

func greatter(values ...int) int {
	if len(values) < 1 {
		return values[0]
	}
	greatter := values[0]
	for _, v := range values[1:] {
		if v > greatter {
			greatter = v
		}
	}
	return greatter
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fib(n uint) uint {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}

func f() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("RECOVERED", x)
		}
	}()
	panic("Panic!")
}

func memoryAddress() {
	x := 2
	fmt.Println(&x)
}

func pointerAtribute() {
	x := new(int)
	fmt.Println(*x)
	*x = 2
	fmt.Println(*x)
}

func newPointer() {
	x := new(int)
	fmt.Println(&x)
}

func square(x *int) {
	*x = *x * *x
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	// a := makeOddGenerator()
	// for i := 0; i < 10 ; i++ {
	// 	fmt.Println(a())
	// }
	// fmt.Println(fib(31))
	// newPointer()
	// x := 2
	// fmt.Println(x)
	// square(&x)
	// fmt.Println(x)
	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println(a, b)
}
