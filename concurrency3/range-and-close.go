package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	channel := make(chan int, 10)
	go fibonacci(cap(channel)*2, channel)
	for i := range channel {
		fmt.Println(i)
	}
}
