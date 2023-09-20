package main

import (
	"fmt"
	"strings"
)

func stringsContains() {
	fmt.Println(strings.Contains("teste", "est"))
}

func stringsCount() {
	fmt.Println(strings.Count("teste", "t"))
}

func stringsHasPrefix() {
	fmt.Println(strings.HasPrefix("teste", "te"))
}

func stringsHasSufix() {
	fmt.Println(strings.HasSuffix("teste", "te"))
}

func stringsIndex() {
	fmt.Println(strings.Index("yellow", "l"))
}

func stringsJoin() {
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
}

func stringsRepeat() {
	fmt.Println(strings.Repeat("-", 10))
}

func stringsReplace() {
	// -1 replace all
	fmt.Println(strings.Replace("aaaaaa", "a", "b", 3))
}

func stringsSplit() {
	fmt.Println(strings.Split("a-b-c-d-e-f", "-"))
}

func stringsToLower() {
	fmt.Println(strings.ToLower("FROM UPPER TO LOWER"))
}

func stringsToUpper() {
	fmt.Println(strings.ToUpper("from lower to upper"))
}

func bytes() {
	a := []byte("string")
	fmt.Println("a: ", a)
	b := string(a)
	fmt.Println("b: ", b)
}

func main() {
	bytes()
}