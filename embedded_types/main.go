package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) talk(msg string) {
	a := fmt.Sprintf("%s Talk: %s", p.name, msg)
	fmt.Println(a)
}

type Android struct {
	Person
	model string
}


func main() {
	a := Android{ Person{"C3P0"}, "C3P0"}
	a.talk("Hello")
}