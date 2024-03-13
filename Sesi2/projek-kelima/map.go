package main

import "fmt"

func main() {
	// map dasar
	// var person map[string]string // Deklarasi

	// person = map[string]string{}

	// person["name"] = "Aurellia"

	// person["age"] = "21"

	// person["address"] = "Jakarta"

	// fmt.Println("name ==>", person["name"])
	// fmt.Println("age ==>", person["age"])
	// fmt.Println("address ==>", person["address"])

	// map dasar lain
	// var person = map[string]string{
	// 	"name":    "Aurellia",
	// 	"age":     "21",
	// 	"address": "Jakarta",
	// }

	// fmt.Println("name:", person["name"])
	// fmt.Println("age:", person["age"])
	// fmt.Println("address:", person["address"])

	// map (looping with map)
	// var person = map[string]string{
	// 	"name":    "Aurellia",
	// 	"age":     "21",
	// 	"address": "Jakarta",
	// }

	// for key, value := range person {
	// 	fmt.Println(key, ":", value)
	// }

	// map (deleting value)
	// var person = map[string]string{
	// 	"name":    "Aurellia",
	// 	"age":     "21",
	// 	"address": "Jakarta",
	// }

	// fmt.Println("Before deleting:", person)

	// delete(person, "age")

	// fmt.Println("After deleting:", person)

	// map (detecting a value)
	// var person = map[string]string{
	// 	"name":    "Aurellia",
	// 	"age":     "21",
	// 	"address": "Jakarta",
	// }

	// value, exist := person["age"]

	// if exist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Value does'nt exist")
	// }

	// delete(person, "age")

	// value, exist = person["age"]

	// if exist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Value has been deleted")
	// }

	// map (combining slice with map)
	var people = []map[string]string{
		{"name": "Aurellia", "age": "21"},
		{"name": "Lulu", "age": "21"},
		{"name": "Zee", "age": "19"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}
}
