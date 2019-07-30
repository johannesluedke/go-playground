package main

import "fmt"
import "time"

func main() {
	fmt.Println("Hello World")
	// Jan 2 15:04:05 2006 MST
	fmt.Println("Now is the " + time.Now().Format("02.01.2006 15:04"))
	/*
		This is a multi-line comment
	*/
	fmt.Println("statement separated by ';'")
	fmt.Println("second")

	var firstVar, secondVar int
	firstVar = 5
	secondVar = 8
	fmt.Println(firstVar * secondVar)

	// weiter mit https://www.tutorialspoint.com/go/go_variables.htm
}
