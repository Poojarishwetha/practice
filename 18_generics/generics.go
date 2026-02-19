// package main

// import "fmt"

// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice(names []string) {
// 	for _, name := range names {
// 		fmt.Println(name)
// 	}
// }

// func main() {
// 	// nums := []int{1, 2, 3}
// 	names := []string{"golang", "typescript"}
// 	printStringSlice(names)
// }

// using generics
// this is for functions
package main

import "fmt"

// func printSlice[T any](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func main() {
// 	num := []int{1, 2, 3, 4, 5}
// 	printSlice(num)

// 	name := []string{"golang", "js", "sql"}
// 	printSlice(name)
// }

//generics can be used in struct also

type stack[T any] struct {
	elements []T
}

func main() {
	myStack := stack[string]{
		elements: []string{"golang"},
	}
	mStack := stack[int]{
		elements: []int{1, 2, 3},
	}
	fmt.Println(myStack)
	fmt.Println(mStack)
}

//comparable-an interface
//multiple generic types can be passed
