package main

import "fmt"

func main() {
	// condition
	// var currentYear = 2021

	// if age := currentYear - 1998; age < 17 {
	// 	fmt.Println("Kamu belum bisa membuat kartu SIM")
	// } else {
	// 	fmt.Println("Kamu sudah boleh membuat kartu SIM")
	// }

	// Switch
	// var score = 8

	// switch score {
	// case 8:
	// 	fmt.Println("perfect")
	// case 7:
	// 	fmt.Println("awesome")
	// default:
	// 	fmt.Println("not bad")
	// }

	// Switch with relational operator
	// var score = 6

	// switch {
	// case score == 8:
	// 	fmt.Println("Perfect")
	// case (score < 8) && (score > 3):
	// 	fmt.Println("not bad")
	// default:
	// 	{
	// 		fmt.Println("study harder")
	// 		fmt.Println("you need to learn more")
	// 	}
	// }

	// switch fallthrough keyword
	//

	var score = 10

	if score > 7 {
		switch score {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score == 5 {
			fmt.Println("not bad")
		} else if score == 3 {
			fmt.Println("you can do it")
			if score == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
