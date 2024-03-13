package main

import (
	"fmt"
)

func main() {
	// slice dasar
	// var buah = []string{"stroberi", "pisang", "apel"}

	// _ = buah

	// make function
	// var buah = make([]string, 3)

	// _ = buah

	// fmt.Printf("%#v\n", buah)

	// slice tanpa append function

	// var buah = make([]string, 3)
	// _ = buah

	// buah[0] = "stroberi"
	// buah[1] = "pisang"
	// buah[2] = "apel"

	// fmt.Printf("%#v\n", buah)

	// slice append function
	// var buah = make([]string, 3)
	// buah = append(buah, "stroberi", "pisang", "apel")

	// fmt.Printf("%#v\n", buah)

	// slice append function with ellipsis
	// var buah1 = []string{"stroberi", "pisang", "apel"}
	// var buah2 = []string{"durian", "mangga", "semangka"}

	// buah1 = append(buah1, buah2...)

	// fmt.Printf("%#v\n", buah1)

	// slice copy function
	// var buah1 = []string{"stroberi", "pisang", "apel"}
	// var buah2 = []string{"durian", "mangga", "semangka"}

	// nn := copy(buah1, buah2)

	// fmt.Println("Buah1 =>", buah1)
	// fmt.Println("Buah2 =>", buah2)
	// fmt.Println("Copy =>", nn)

	// slice (slicing)
	// var buah1 = []string{"stroberi", "pisang", "apel", "jeruk", "anggur"}

	// var buah2 = buah1[1:4]
	// fmt.Printf("%#v\n", buah2)

	// var buah3 = buah1[0:]
	// fmt.Printf("%#v\n", buah3)

	// var buah4 = buah1[:3]
	// fmt.Printf("%#v\n", buah4)

	// var buah5 = buah1[:] //sama dengan buah1[:len(buah1)]
	// fmt.Printf("%#v\n", buah5)

	// slice (combining slicing and append)
	// var buah1 = []string{"stroberi", "pisang", "apel", "jeruk", "anggur"}

	// buah1 = append(buah1[:3], "mangga")

	// fmt.Printf("%#v\n", buah1)

	// slice (backing array)
	// var buah1 = []string{"stroberi", "pisang", "apel", "jeruk", "anggur"}

	// var buah2 = buah1[2:4]

	// buah2[0] = "mangga"

	// fmt.Println("Buah1 =>", buah1)
	// fmt.Println("Buah2 =>", buah2)

	// slice (cap function)
	// var buah1 = []string{"stroberi", "pisang", "apel", "jeruk"}

	// fmt.Println("Buah1 cap:", cap(buah1)) // 4
	// fmt.Println("Buah1 len:", len(buah1)) // 4

	// fmt.Println(strings.Repeat("#", 20))

	// var buah2 = buah1[0:3]

	// fmt.Println("Buah2 cap:", cap(buah2)) // 4
	// fmt.Println("Buah2 len:", len(buah2)) // 3

	// fmt.Println(strings.Repeat("#", 20))

	// var buah3 = buah1[1:]

	// fmt.Println("Buah3 cap:", cap(buah3)) // 3
	// fmt.Println("Buah3 len:", len(buah3)) // 3

	// slice (creating a new backing array)
	cars := []string{"Ford", "Honda", "Audi", "BMW"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Toyota"
	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)

}
