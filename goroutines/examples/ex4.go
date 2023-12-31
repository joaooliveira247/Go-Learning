// Select: an switch to channels

package examples

import (
	"fmt"
	"time"
)

func RunEx4() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func()  {
		for {
			c1 <- "From 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func()  {
		for {
			c2 <- "From 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func ()  {
		for {
			select {
			case msg1 := <- c1:
				fmt.Println(msg1)
			case msg2 := <- c2:
				fmt.Println(msg2)
			case <- time.After(time.Second):
				fmt.Println("timeout")
			default:
				fmt.Println("not read")
			}
			
		}
	}()

	var input string
	fmt.Scanln(&input)
}