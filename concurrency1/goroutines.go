package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		say("tic")
		wg.Done()
	}()
	go func() {
		say("toc")
		wg.Done()
	}()
	go func() {
		fmt.Println("yay!")
		wg.Done()
	}()
	wg.Wait()
}
