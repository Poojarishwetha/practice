package main

import "fmt"

//range -> keyword which is used to iterate over data strucutres

func main() {
	nums := []int{6, 7, 8}
	sum := 0
	for _, num := range nums {
		sum += num

	}
	println(sum)

	//********************
	n := []int{6, 7, 8}

	for i, m := range n { //i is index
		fmt.Println(m, i)

	}
	//***************************************

	m := map[string]string{"fname": "john", "lname": "doe"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	//we can use range upon strings also
	for i, c := range "golang" { //unicode code point rune
		fmt.Println(i, c)
		fmt.Println(i, string(c))
	}
}
