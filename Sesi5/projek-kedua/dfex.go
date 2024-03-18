package main

// func main() {
// **defer sederhana
// 	defer fmt.Println("hello world")
// 	fmt.Println("lutfi")
// 	f()
// }

// func f() {
// 	defer fmt.Println("farhan")
// 	fmt.Println("hakim")
// }

// **latihan1
// func main() {
// 	for i := 0; i < 10; i++ {
// 		defer fmt.Println(i)
// 	}
// }

// **latihan2
// func main() {
// 	var a int = 5
// 	defer fmt.Println(a)

// 	a = 10
// 	fmt.Println("ini a", a)
// }

// **latihan3
// func main() {
// 	var a int = 5
// defer func() {
// 	fmt.Println("ini dari defer", a)
// }() //**yang muncul 10

// 	defer func(x int) {
// 		fmt.Println("ini dari defer", x)
// 	}(a) //**yang muncul 5

// 	a = 10
// 	fmt.Println("ini a", a)
// }

// **latihan4
// func main() {
// 	var a int = 5
// 	defer func() {
// 		fmt.Println("ini dari defer", a)
// 	}()

// 	a = 123
// 	fmt.Println("ini a", a)

// 	for i := 0; i < 10; i++ {
// 		defer func() {
// 			fmt.Println(i)
// 		}()
// 	}
// }

// **latihan5 jika konek ke database
// func main() {
// 	var con = ConnectDatabase()
// 	defer con.Close()

// 	if blabla {
// 		return
// 	}

// 	if blablabla {
// 		return
// 	} else {

// 	}
// }
