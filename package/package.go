package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(1))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("Pi: ", math.Pi)
	fmt.Println(add(42, int(math.Ceil(23.330))))
}
