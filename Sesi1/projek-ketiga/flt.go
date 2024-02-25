package main

import "fmt"

func main() {
	var decimalnumber float32 = 3.63

	fmt.Printf("decimal number: %f\n", decimalnumber)   //memformat desimal menjadi 6 digit angka
	fmt.Printf("decimal number: %.3f\n", decimalnumber) //mengecilkan desimal
}
