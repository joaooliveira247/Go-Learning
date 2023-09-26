package examples

import (
	"fmt"
	"math/rand"
	"time"
)

func f2(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func RunEx2() {
	fmt.Println("Runing example 2")
	for i := 0; i < 10; i++ {
		go f2(i)
	}
	var input string
	fmt.Scanln(&input)
}
