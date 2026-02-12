package main

import (
	"fmt"
	"maps"
)

// maps(associative datastructures)->hash tables, objects , dictionaries in other languages. Maps are used to store key-value pairs. Maps are declared using the make function
func main() {
	//creating map using make function
	m := make(map[string]string)

	//adding key-value pairs to the map
	m["name"] = "golang"
	m["area"] = "backend"

	//get an element
	fmt.Println(m["name"], m["area"])
	fmt.Println(m["phone"]) //if key value does not exists in the map, then it returns zero value

	n := make(map[string]int)
	n["age"] = 30
	fmt.Println(n["age"])
	fmt.Println(n["phone"])

	//to get the length of the map
	fmt.Println(len(m))

	delete(m, "area")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)
	//************************************************************

	//important feature of maps
	p := map[string]int{"price": 100, "phone": 5}
	// fmt.Println(p)

	//to check if the key exists in the map or not, we can use the comma ok idiom. It returns two values, the first value is the value of the key and the second value is a boolean value which indicates whether the key exists in the map or not.

	s, ok := p["price"]
	fmt.Println(s)
	if ok { //can also use !ok to check if the key does not exists in the map
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	//***************************************************

	//to check both the maps are equal or not
	m1 := map[string]int{"price": 100, "phone": 5}
	m2 := map[string]int{"price": 100, "phone": 6}
	fmt.Println(maps.Equal(m1, m2))
}
