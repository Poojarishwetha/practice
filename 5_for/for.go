package main

import "fmt"

//for - > only construct in go for looping
func main() {
	//while loop.in go we dont have while loop but we can achieve the same functionality using for loop
	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//infinite loop
	// for {
	// 	println(1)
	// }

	//classic for loop
	// for i := 0; i <= 6; i++ {
	// fmt.Println(i)
	// break //break statement is used to exit the loop when a certain condition is met. In this case, we are breaking the loop when i is equal to 3.
	// if i == 3 {
	// 	continue
	// }
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	println(i)
	// }

	//range loop - > used to iterate over arrays, slices, maps, strings, channels etc. It returns the index and value of the element in each iteration.
	for i := range 10 {
		fmt.Println(i)
	}

}
