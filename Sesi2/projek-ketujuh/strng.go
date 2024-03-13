package main

import (
	"fmt"
)

func main() {
	// looping over string (byte-by-byte)
	// city := "Jakarta"

	// for i := 0; i < len(city); i++ {
	// 	fmt.Printf("Index: %d, byte: %d\n", i, city[i])
	// }

	// var city []byte = []byte{74, 97, 107, 97, 114, 116, 97}

	// fmt.Println(string(city))

	// looping over string (rune-by-rune)
	// str1 := "makan"

	// str2 := "mânca"

	// fmt.Printf("str1 byte length: %d\n", len(str1))
	// fmt.Printf("str2 byte length: %d\n", len(str2))

	// str1 := "makan"

	// str2 := "mânca"

	// fmt.Printf("str1 rune length: %d\n", utf8.RuneCountInString(str1))
	// fmt.Printf("str2 rune length: %d\n", utf8.RuneCountInString(str2))

	str := "mânca"

	for i, s := range str {
		fmt.Printf("Index: %d, string: %s\n", i, string(s))
	}
}
