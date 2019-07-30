package main

import (
	"fmt"
	"math"
)

var i, j = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	k := 3
	fmt.Println(i, j, c, python, java)
	fmt.Println(k)

	var i int
	var f float32
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var ii = -42
	var ff = float64(ii)
	// this does not go well:
	u := uint(ff)
	fmt.Println(u)

	var x, y int = 3, 4
	var fff = math.Sqrt(float64(x*x + y*y))
	var zz = uint(fff)
	fmt.Println(x, y, zz)

	v := -42
	fmt.Printf("v is of type %T\n", v)
}
