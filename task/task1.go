package main

import "fmt"

func add(new []int) {
	new = append(new, 2)
	new = append(new, 5)
	fmt.Println(new)
}

func main() {
	new := make([]int, 2, 5)

	new[0] = 1
	new[1] = 2
	add(new)

}
