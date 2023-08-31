package main

import "fmt"

func declare_var() {
	var x string = "Hello, World!"
	fmt.Println(x)
}

func declare_new_empty_var() {
	var x string
	x = "Hello, World!"
	fmt.Println(x)
}

func var_assignment() {
	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)
}

func concat_assignment() {
	var x string
	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)
}

func comparacion() {
	var x string = "Hello"
	var y string = "World"
	fmt.Println(x == y)
	y = "Hello"
	fmt.Println(x == y)
}

func first_var_assignment() {
	x := "Hello World!"
	fmt.Println(x)
}

func var_name_code_style() {
	// Go uses the shit of camelCase
	dogName := "max"
	fmt.Println("My dog's name is", dogName)
}

func consts() {
	const hello = "Hello"
	fmt.Println(hello)
	const typedHello string = "Hello typed"
	fmt.Println(typedHello)
}

func multiple_var_assignment() {
	var (
		a = "2"
		b = 2
		c = 5.0
	)
	fmt.Println(a, b, c)
}

func double() {
	fmt.Println("Type a number.")
	var num float64
	fmt.Scanf("%f", &num)

	output := num * 2

	fmt.Println(output)

}

func main() {
	double()
}
