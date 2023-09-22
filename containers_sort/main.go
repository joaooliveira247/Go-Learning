package main

import (
	"container/list"
	"fmt"
	"sort"
	"strings"
)

func containerList() {
	// list accept multiple types.
	// list.New()
	var x list.List
	x.PushBack(1)
	x.PushBack("Two")
	x.PushBack(false)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

type Person struct {
	Name string
	Age int
}

type ByName []Person

type ByAge []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps ByAge) Len() int {
	return len(ps)
}

func (ps ByAge) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}

func (ps ByAge) Swap(i,j int) {
	ps[i], ps[j] = ps[j], ps[i] 
}

func main() {
	kids := []Person{
		{"Jill", 8},
		{"Jane", 9},
		{"Juan", 6},
		{"John", 22},
	}
	fmt.Println(strings.Repeat("-", 11))
	fmt.Println(kids)
	sort.Sort(ByName(kids))
	fmt.Println(kids)
	fmt.Println(strings.Repeat("-", 11))
	fmt.Println(kids)
	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}
