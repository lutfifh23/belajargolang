package main

import (
	"fmt"
	"runtime"
	"time"
)

// goroutine dasar
// func main() {
// 	go goroutine()
// }

// func goroutine() {
// 	fmt.Println("Hello World")
// }

// goroutine (asynchronous process #1&2)
func main() {
	// 	fmt.Println("main execution started")

	// 	go firstProcess(8)

	// 	seconProcess(8)

	// 	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	// 	fmt.Println("main execution ended")
	// }

	// func firstProcess(index int) {
	// 	fmt.Println("first process func started")
	// 	for i := 1; i <= index; i++ {
	// 		fmt.Println("i=", i)
	// 	}
	// 	fmt.Println("first process func ended")
	// }

	// func seconProcess(index int) {
	// 	fmt.Println("second process func started")
	// 	for j := 1; j <= index; j++ {
	// 		fmt.Println("j=", j)
	// 	}
	// 	fmt.Println("second process func ended")

	fmt.Println("main execution started")

	go firstProcess(8)

	secondProcess(8)

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	time.Sleep(time.Second * 2)

	fmt.Println("main execution ended")
}

func firstProcess(index int) {
	fmt.Println("first process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("first process func ended")
}

func secondProcess(index int) {
	fmt.Println("second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("second process func ended")
}
