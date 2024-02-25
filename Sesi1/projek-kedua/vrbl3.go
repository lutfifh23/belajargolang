// variable tanpa type data tapi menampilkan tipe data menggunakan printf
package main

import "fmt"

func main() {
	var name = "Aurell"
	var age = 21

	fmt.Printf("%T, %T", name, age)
}
