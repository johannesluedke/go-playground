package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func printInfo() int {
	fmt.Println("some info")
	return 1
}

func main() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}

	i := rand.Intn(2)
	switch i {
	case 0:
	case printInfo():
		printInfo()
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("too far away")
	}

	//switch with no condition is the same as `switch true`
	fmt.Println("===")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
