package main

import "fmt"

// define a struct
type Person struct {
	Name string
	Age  int
}

func main() {
	// create a struct variable
	p := Person{
		Name: "Ravi",
		Age:  25,
	}

	// access struct fields
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
