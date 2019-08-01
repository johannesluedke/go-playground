package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //send sum to channel c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from channel

	fmt.Println(x, y, x+y)

	// buffered channel
	bufferedChannel := make(chan int, 2)
	bufferedChannel <- 1
	bufferedChannel <- 2
	fmt.Println(<-bufferedChannel, <-bufferedChannel)
}
