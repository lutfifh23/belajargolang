package main

import (
	"fmt"
)

// struct dasar
// type Employee struct {
// 	name     string
// 	age      int
// 	division string
// }

// func main() {

// }

// struct (Giving value to struct)
// type Employee struct {
// 	name     string
// 	age      int
// 	division string
// }

// func main() {
// 	var employee Employee

// 	employee.name = "Aurellia"

// 	employee.age = 21

// 	employee.division = "IT"

// 	fmt.Println(employee.name)
// 	fmt.Println(employee.age)
// 	fmt.Println(employee.division)
// }

// struct (initializing struct)
// type Employee struct {
// 	name     string
// 	age      int
// 	division string
// }

// func main() {
// 	var employee1 = Employee{}
// 	employee1.name = "Aurellia"
// 	employee1.age = 21
// 	employee1.division = "IT"

// 	var employee2 = Employee{name: "Lulu", age: 21, division: "IT"}

// 	fmt.Printf("Employee1: %+v\n", employee1)
// 	fmt.Printf("Employee2: %+v\n", employee2)
// }

// struct (pointer to struct)
// type Employee struct {
// 	name     string
// 	age      int
// 	division string
// }

// func main() {
// 	var employee1 = Employee{name: "Aurellia", age: 21, division: "IT"}

// 	var employee2 *Employee = &employee1

// 	fmt.Println("Employee1 name:", employee1.name)
// 	fmt.Println("Employee2 name:", employee2.name)

// 	fmt.Println(strings.Repeat("#", 20))

// 	employee2.name = "Lulu"

// 	fmt.Println("Employee1 name:", employee1.name)
// 	fmt.Println("Employee2 name:", employee2.name)
// }

// struct (embedded struct)
// type Person struct {
// 	name string
// 	age  int
// }

// type Employee struct {
// 	division string
// 	person   Person
// }

// func main() {
// 	var employee1 = Employee{}

// 	employee1.person.name = "Aurellia"
// 	employee1.person.age = 21
// 	employee1.division = "IT"

// 	fmt.Printf("%+v", employee1)
// }

// struct (anonymous struct)
// type Person struct {
// 	name string
// 	age  int
// }

// func main() {
// 	// anonymous struct tanpa pengisian field
// 	var employee1 = struct {
// 		person   Person
// 		division string
// 	}{}
// 	employee1.person = Person{name: "Aurellia", age: 21}
// 	employee1.division = "IT"

// 	// anonymous struct dengan pengisian field
// 	var employee2 = struct {
// 		person   Person
// 		division string
// 	}{
// 		person:   Person{name: "Lulu", age: 21},
// 		division: "IT",
// 	}

// 	fmt.Printf("Employee1: %+v\n", employee1)
// 	fmt.Printf("Employee2: %+v\n", employee2)
// }

// struct (slice of struct)
// type Person struct {
// 	name string
// 	age  int
// }

// func main() {
// 	var people = []Person{
// 		{name: "Aurellia", age: 21},
// 		{name: "Lulu", age: 21},
// 		{name: "Flora", age: 18},
// 	}

// 	for _, v := range people {
// 		fmt.Printf("%+v\n", v)
// 	}
// }

// struct (slice of anonymous struct)
type Person struct {
	name string
	age  int
}

func main() {
	var employee = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Aurellia", age: 21}, division: "IT"},
		{person: Person{name: "Lulu", age: 21}, division: "IT"},
		{person: Person{name: "Flora", age: 18}, division: "HR"},
	}

	for _, v := range employee {
		fmt.Printf("%+v\n", v)
	}

}
