package main

import "fmt"

// **channel dasar
// func main() {
// 	// var a int = 5

// 	// var c chan int = make(chan int) //unbuffer channel
// 	var c chan int = make(chan int, 4) //buffer channel
// 	go someFunc(c)
// 	// c <- 1
// 	fmt.Println("hello world1")
// 	// fmt.Println(<-c)
// 	time.Sleep(5 * time.Second)
// 	fmt.Println(<-c)
// 	// fmt.Println("hello world2")
// }

// func someFunc(iniChannelPalingGokil chan int) {
// 	fmt.Println("hello world3")
// 	// fmt.Println(<-iniChannelPalingGokil)
// 	iniChannelPalingGokil <- 123
// 	// iniChannelPalingGokil <- 456
// 	// iniChannelPalingGokil <- 789 //ga bakal muncul karena yang manggilnya cuman 2
// 	fmt.Println("hello world4") //tidak akan muncul apabila data lebih dari 4
// }

// **latihan

// func main() {
// 	c := make(chan int)

// 	go channel(c)
// 	c <- 10
// 	fmt.Println(<-c)

// 	go parameter(c, 11)
// 	fmt.Println(<-c)
// }

// func channel(c chan int) {
// 	input := <-c
// 	c <- 10 + input
// }

// func parameter(c chan int, n int) {
// 	input := n
// 	c <- 10 + input
// }

// **channel select
func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go f1(c1)
	go f2(c2)

	for i := 0; i < 2; i++ {
		select {
		case <-c1:
			fmt.Println("from c1")
		case <-c2:
			fmt.Println("from c2")
		}
	}
}

func f1(c chan int) {
	c <- 1
}

func f2(c chan int) {
	c <- 2
}
