// operator logika
package main

import "fmt"

func main() {
	var right = true
	var wrong = false

	var WrongandRight = wrong && right
	fmt.Printf("wrong && right \t(%t)\n", WrongandRight)

	var WrongorRight = wrong || right
	fmt.Printf("wrong || right \t(%t)\n", WrongorRight)

	var WrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t)\n", WrongReverse)
}
