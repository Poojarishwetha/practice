package main

import "fmt"

//return type is function but not recieveing anything as parameter
func counter() func() int {
	var count int = 0

	return func() int {
		count++
		return count
	}

}

func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
}
