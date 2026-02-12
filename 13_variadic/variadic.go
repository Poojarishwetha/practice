package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// fmt.Println(1, 2, 3, 4, "hello", true) //can pass n number of parameters

	// result := sum(3, 4, 5, 6)
	// fmt.Println(result)

	nums := []int{1, 9, 34, 78}
	result := sum(nums...)
	fmt.Println(result)
}
