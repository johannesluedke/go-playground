package main

import (
	"fmt"
	"math"
)

func ownSqrt(x float64) float64 {
	z := x / 2.0 // initial guess
	i := 0
	// add some limit to prevent endless blocking
	for j := 0; j < 10000; j++ {
		i++
		changeZ := (z*z - x) / (2 * z)
		z -= changeZ
		fmt.Println(math.Abs(changeZ))
		fmt.Println(math.Pow(10, -6))
		if math.Abs(changeZ) < math.Pow(10, -13) {
			fmt.Println("iterated", i, "times")
			return z
		}
	}
	return z
}

func main() {
	value := 2.0
	ownResult := ownSqrt(value)
	acutalResult := math.Sqrt(value)
	fmt.Println("====")
	fmt.Println("value is: ", ownResult)
	fmt.Println("value should be: ", acutalResult)
	fmt.Println("difference: ", ownResult-acutalResult)

}
