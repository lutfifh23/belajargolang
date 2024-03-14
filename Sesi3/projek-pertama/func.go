package main

import (
	"fmt"
	"strings"
)

func main() {
	// function dasar
	// 	greet("Aurellia", 21)
	// }

	// func greet(name string, age int8) {
	// 	fmt.Printf("Hello there! My name is %s and I am %d years old.\n", name, age)

	// function dasar lain
	// 	greet("Aurellia", "Jakarta")
	// }
	// func greet(name, address string) {
	// 	fmt.Println("Hello there! My name is", name)
	// 	fmt.Println("I live in", address)

	// function (return)
	// 	var names = []string{"Aurellia", "Lulu"}
	// 	var printMsg = greet("Haii", names)

	// 	fmt.Println(printMsg)
	// }

	// func greet(msg string, names []string) string {
	// 	var joinStr = strings.Join(names, ", ")

	// 	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	// 	return result

	// function (returning multiple values)
	// 	var diameter float64 = 15

	// 	var area, circumference = calculate(diameter)

	// 	fmt.Println("Area:", area)
	// 	fmt.Println("Circumference:", circumference)
	// }

	// func calculate(d float64) (float64, float64) {
	// 	// menghitung luas
	// 	var area float64 = math.Pi * math.Pow(d/2, 2)
	// 	// menghitung keliling
	// 	var circumference = math.Pi * d

	// 	return area, circumference

	// function (predefined return value)
	// 	var diameter float64 = 15

	// 	var area, circumference = calculate(diameter)

	// 	fmt.Println("Area:", area)
	// 	fmt.Println("Circumference:", circumference)
	// }

	// func calculate(d float64) (area float64, circumference float64) {
	// 	// menghitung luas
	// 	area = math.Pi * math.Pow(d/2, 2)
	// 	// menghitung keliling
	// 	circumference = math.Pi * d

	// 	return

	// function (variadic function #1)
	// 	studentList := print("Aurellia", "Lulu", "Zee", "Adel", "Oniel")

	// 	fmt.Printf("%v", studentList)
	// }

	// func print(names ...string) []map[string]string {
	// 	var result []map[string]string

	// 	for i, v := range names {
	// 		key := fmt.Sprintf("student%d", i+1)
	// 		temp := map[string]string{
	// 			key: v,
	// 		}
	// 		result = append(result, temp)
	// 	}
	// 	return result

	// function (variadic function #2)
	// 	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// 	result := sum(numberList...)

	// 	fmt.Println("Result:", result)
	// }

	// func sum(numbers ...int) int {
	// 	total := 0

	// 	for _, v := range numbers {
	// 		total += v
	// 	}
	// 	return total

	// function (variadic function #3)
	profile("Aurellia", "Pasta", "Ayam Geprek", "Es Jeruk", "Kopi Hitam")
}

func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ", ")

	fmt.Println("Hello there! My name is", name)
	fmt.Println("My favorite foods are", mergeFavFoods)
}
