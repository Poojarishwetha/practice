package main

import "fmt"

func main() {
	// age := 12
	// if age >= 18 {
	// 	fmt.Println("eligible to vote")
	// } else {
	// 	fmt.Println("not eligible to vote")
	// }

	// age := 4
	// if age >= 18 {
	// 	fmt.Println("adult")
	// } else if age >= 12 {
	// 	fmt.Println("teenager")
	// } else {
	// 	fmt.Println("child")
	// }

	//can use logical operators in if else statements
	// role := "admin"
	// hasPermission := true
	// if role == "admin" || hasPermission {
	// 	fmt.Println("yes")
	// }

	// if role == "admin" && hasPermission {
	// 	fmt.Println("no")
	// }

	//we can declare variable inside if constructor
	if age := 10; age >= 18 {
		fmt.Println("adult")
	} else if age >= 12 {
		fmt.Println("teenager")
	} else {
		fmt.Println("child")
	}



	//go doesnt have ternary operator like other languages, we can use if else statements instead of ternary operator
}
