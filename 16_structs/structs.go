package main

import "fmt"

// define a struct
type Person struct {
	Name string
	Age  int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	H_no  int
	Area  string
	State string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	// create a struct variable
	//1st method
	// p := Person{
	// 	Name: "Ravi",
	// 	Age:  25,
	// }

	//2nd method
	// var p Person
	// p.Name = "shwetha"
	// p.Age = 22

	//new keyword
	// var p = new(Person)
	// p.Name = "simran"
	// p.Age = 22

	// access struct fields
	// fmt.Println("Name:", p.Name)
	// fmt.Println("Age:", p.Age)
	// fmt.Println("details:", *p)

	var emp Employee
	emp.Person_Details = Person{
		Name: "raj",
		Age:  22,
	}

	emp.Person_Contact.Email = "abc@gmail.com"
	emp.Person_Contact.Phone = "9876899887"

	emp.Person_Address = Address{
		H_no:  23,
		Area:  "kjh",
		State: "qwertyui",
	}
	fmt.Println(emp.Person_Address)
	fmt.Println(emp.Person_Details)

}
