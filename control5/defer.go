package main

import "fmt"

func subfunction() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main() {
	defer fmt.Println("world")

	fmt.Println("Hello")

	subfunction()
}
