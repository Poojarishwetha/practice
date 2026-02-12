package main

import (
	"fmt"
)

// when we have multiple conditions to check, we can use switch statement instead of if-else statement. Switch statement is more readable and easier to understand than if-else statement. Switch statement is used to check the value of a variable against multiple cases. Switch statement can be used with any data type.
func main() {
	//simple switch
	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other number")
	// }

	//multiple condition switch
	// switch time.Now().Hour() {
	// case 0, 1, 2, 3, 4, 5, 6:
	// 	fmt.Println("good morning")
	// default:
	// 	fmt.Println("good evening")
	// }

	//type switch->poweful switch statement
	WhoAmI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("it is an integer")
		case string:
			fmt.Printf("it is a string")
		case bool:
			fmt.Printf("it is a boolean")
		default:
			fmt.Printf("unknown type %T", v)
		}
	}
	WhoAmI(55.7)
}
