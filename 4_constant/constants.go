package main

import "fmt"

//constants - once declared, the value of a constant cannot be changed. Constants are declared using the const keyword.
//can declare constants outside the function also
//we cannot use := operator to declare constants, we have to use const keyword to declare constants
var age = 30

func main() {
	const name string = "golang"
	fmt.Println(name)
	fmt.Println(age)

	const (
		country = "India"
		city    = "Delhi"
		gender  = "male"
	)
	fmt.Println(gender, country)

}
