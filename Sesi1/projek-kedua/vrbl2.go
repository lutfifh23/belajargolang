package main

import "fmt"

func main() {
	var name string
	var age int

	name = "Aurellia"
	age = 21

	fmt.Println("Ini adalah namanya ==>", name)
	fmt.Println("Ini adalah umurnya ==>", age)
}

// contoh error
// package main

// import "fmt"

// func main() {
// 	var name string
// 	var age int

// 	name = 21 *harusnya diisi sama string
// 	age = "Aurellia" *harusnya diisi sama integer

// 	fmt.Println("Ini adalah namanya ==>", name)
// 	fmt.Println("Ini adalah umurnya ==>", age)
// }
