package main

// func add(a int, b int) int {
func add(a, b int) int {
	return a + b
}

func getLanguages() (string, string, string) {
	return "golang", "javascript", "c"

}

//using function we can return new function also
func processIt(fn func(a int) int) {
	fn(1)
}

func main() {
	// result := add(1, 2)
	// fmt.Println(result)
	// fmt.Println(add(1, 5))

	// fmt.Println(getLanguages())  or
	// lang1, lang2, lang3 := getLanguages()
	// fmt.Println(lang1, lang2, lang3)

	fn := func(a int) int {
		return 2
	}

	processIt(fn)
}
