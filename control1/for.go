package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum of 0 to 10: ", sum)

	// also works like this:
	sum2 := 1
	for sum2 < 1000 { // equals C's `while` loop
		sum2 += sum2
		fmt.Println(sum2)
	}
	fmt.Println("second sum", sum2)
}
