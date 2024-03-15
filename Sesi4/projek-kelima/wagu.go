package main

import (
	"fmt"
	"sync"
)

func main() {
	fruits := []string{"stroberi", "pisang", "apel", "anggur"}

	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1)
		go printfruit(index, fruit, &wg)
	}

	wg.Wait()
}

func printfruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("index => %d, fruit => %s\n", index, fruit)
	wg.Done()
}
