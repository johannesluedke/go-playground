package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

func main() {
	fmt.Println("My favorite number is", rand.Intn(1))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("Pi: ", math.Pi)
	fmt.Println(add(42, int(math.Ceil(23.330))))

	a, b := swap("good", "morning")
	fmt.Println(a, b)
	fmt.Println(swap("first", "then second"))
	fmt.Println(split(9))

	var i int
	fmt.Println(i, c, python, java)
}
