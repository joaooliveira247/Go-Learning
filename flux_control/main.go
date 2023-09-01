package main

import "fmt"

func fluxFor() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}

func anotherFor() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func fluxIf() {
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "| even")
		} else if i % 3 == 0 {
			fmt.Println(i, "| divisible by 3")
		} else {
			fmt.Println(i, "| odd")
		}
	}
}

func fluxSwitch(i uint8) {
	switch i {
	case 0: fmt.Println("Zero")
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	case 3: fmt.Println("Three")
	case 4: fmt.Println("Four")
	case 5: fmt.Println("Five")
	default: fmt.Println("Unknow number")
	}
}

func ex_1() {
	i := 10

	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
}

func ex_2() {
	for i:= 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		} 
	}
}

func ex_3() {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	ex_3()
}