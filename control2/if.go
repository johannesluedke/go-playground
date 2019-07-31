package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// returns x ** n < limit ? x**n : limit
func pow(x, n, limit float64) float64 {
	if value := math.Pow(x, n); value < limit {
		return value
	}
	return limit
}

func main() {
	// immediately called function
	func() {
		fmt.Println("sqrt 2: ", sqrt(2), "sqrt -4:", sqrt(-4))
		fmt.Println(
			pow(3, 2, 1000),
			pow(3, 3, 20),
		)
	}()
}
