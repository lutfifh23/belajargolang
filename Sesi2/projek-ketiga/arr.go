package main

import (
	"fmt"
)

func main() {
	//array dasar
	// var angka [4]int
	// angka = [4]int{1, 2, 3, 4}

	// var strings = [3]string{"aurellia", "lulu", "zee"} // array type

	// fmt.Printf("%#v\n", angka)
	// fmt.Printf("%#v\n", strings)

	//array modify element through index
	// var fruits = [3]string{"stroberi", "pisang", "apel"} // array
	// fruits[0] = "stroberry"
	// fruits[1] = "banana"
	// fruits[2] = "apple"

	// fmt.Printf("%#v\n", fruits)

	//array (loop through elements)
	// var fruits = [3]string{"stroberi", "pisang", "apel"}
	// //Cara Pertama
	// for i, v := range fruits {
	// 	fmt.Printf("Index: %d, Value: %s\n", i, v)
	// }

	// fmt.Println(strings.Repeat("#", 25))

	// //Cara Kedua
	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("Index: %d, Value: %s\n", i, fruits[i])
	// }

	//array (multidimensional array)
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
