package main

import "fmt"

func run1() {
	panic("PANIC!!")
	str := recover() // never be executed
	fmt.Println(str)
}

func run2() {
	defer func() {
		str := recover()
		fmt.Println(str, "recover")
	}()
	panic("PANIC!!")
}

func main() {
	run2()
}