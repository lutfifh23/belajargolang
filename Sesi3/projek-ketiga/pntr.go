package main

import (
	"fmt"
)

func main() {
	// pointer sederhana
	// var firstNumber *int
	// var secondNumber *int
	// _, _ = firstNumber, secondNumber

	// pointer (memory address)
	// var firstNumber int = 4
	// var secondNumber *int = &firstNumber

	// fmt.Println("firstNumber (value):", firstNumber)
	// fmt.Println("firstNumber (address):", &firstNumber)

	// fmt.Println("secondNumber (value):", *secondNumber)
	// fmt.Println("secondNumber (address):", secondNumber)

	// pointer (change value through pointer)
	// var firstPerson string = "Lulu"
	// var secondPerson *string = &firstPerson

	// fmt.Println("firstPerson (value):", firstPerson)
	// fmt.Println("firstPerson (address):", &firstPerson)
	// fmt.Println("secondPerson (value):", *secondPerson)
	// fmt.Println("secondPerson (address):", secondPerson)

	// fmt.Println("\n", strings.Repeat("#", 30), "\n")

	// *secondPerson = "Flora"

	// fmt.Println("firstPerson (value):", firstPerson)
	// fmt.Println("firstPerson (address):", &firstPerson)
	// fmt.Println("secondPerson (value):", *secondPerson)
	// fmt.Println("secondPerson (address):", secondPerson)

	// pointer (pointer as a parameter)
	var a int = 10

	fmt.Println("Before:", a)

	changeValue(&a)

	fmt.Println("After:", a)
}

func changeValue(number *int) {
	*number = 20

}
