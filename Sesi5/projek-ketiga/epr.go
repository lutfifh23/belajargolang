package main

import "fmt"

// **error sederhana
// func main() {
// 	var a int = 5

// 	var b error = errors.New("ini error")

// 	fmt.Println(a, b)
// }

// **latihan error
// func main() {
// 	var a int = 5
// 	value, err := foo()
// 	if err != nil {
// 		return
// 	}
// }

// func foo() (int, error) {

// }

// **panic sederhana
// func main() {
// 	panic("lapar")
// }

// **recover sederhana
func main() {
	defer catch()
	var a int
	fmt.Println(5 / a)
}

func catch() {
	r := recover()
	fmt.Println(r)
	fmt.Println("halo ini lanjut")
}
