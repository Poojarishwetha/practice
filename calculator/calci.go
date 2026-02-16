package main

import "fmt"

func main() {
	var a, b float64
	var op string

	fmt.Println("enter no")
	fmt.Scan(&a)

	for {

		fmt.Println("enter operator +,-,*,/,=")
		fmt.Scan(&op)
		if op == "=" {
			break
		}
		fmt.Println("enter no")
		fmt.Scan(&b)

		switch op {
		case "+":

			a += b
			fmt.Println(a)
		case "-":
			a = a - b
			fmt.Println(a)
		case "*":
			a = a * b
			fmt.Println(a)
		case "/":
			if b == 0 {
				fmt.Println("cannot divide")
			} else {
				a = a / b
				fmt.Println(a)
			}

		default:
			fmt.Println("invalid operator")
		}
	}

}
