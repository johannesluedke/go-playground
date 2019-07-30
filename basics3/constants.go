package main

import "fmt"

// Pi this is Pi! 🐈
const Pi = 3.14

const (
	// Big is huge
	Big = 1 << 100
	// Small is small
	Small = Big >> 99
)

func needInt(x int) int { return 10*x + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const World = "🐈"
	fmt.Println(World)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
