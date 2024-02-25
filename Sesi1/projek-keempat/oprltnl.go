// operator relational
package main

import "fmt"

func main() {
	var firstCondition bool = 2 < 3
	var secondCondition bool = "joy" == "Joy"
	var thirdCondition bool = 10 != 2.3
	var fourthCondition bool = 11 <= 11

	fmt.Println("first condition:", firstCondition)
	fmt.Println("second condition:", secondCondition)
	fmt.Println("third condition:", thirdCondition)
	fmt.Println("fourth condition:", fourthCondition)
}
