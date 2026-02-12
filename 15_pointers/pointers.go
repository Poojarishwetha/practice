package main

import "fmt"

func update(n *int) {
	//derefernce
	*n = 20 //Go to the box that n is pointing to and change its value to 20
}

func main() {
	x := 10
	update(&x) //You give the address of x to the function update.So now the function knows where x lives, not just its value
	fmt.Println(x)
}

//& = “where is it?”     * = “what is there?”
//The program uses pointers to pass the address of a variable to a function, allowing the function to modify the original value.
