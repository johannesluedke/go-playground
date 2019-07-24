package main

import "fmt"
import "time"

func main() {
	fmt.Println("Hello World")
	// Jan 2 15:04:05 2006 MST
	fmt.Println("Now is the " + time.Now().Format("02.01.2006 15:04"))
}
