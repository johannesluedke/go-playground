package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("OK I stop :)")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	var input string
	go func() {
		fmt.Println("press enter if you have seen enough")
		fmt.Scanln(&input)
		quit <- 42
	}()
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			fmt.Println(<-c)
		}
	}()
	fibonacci(c, quit)
}
