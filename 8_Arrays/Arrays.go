package main

import "fmt"

//arrays->numbered sequence of specific length
func main() {
	//zeroed values of array
	//int->0 , string -> "" , bool ->false
	// var nums [5]int //create array
	// to insert elements
	// nums[0] = 1
	// nums[1] = 2
	// nums[2] = 3
	// nums[3] = 4
	// fmt.Println(nums[0])->get element
	// fmt.Println(nums)
	//array length
	// fmt.Println(len(nums))

	// var vals [4]string
	// vals[0] = "hello"
	// fmt.Println(vals)

	// var sample [3]bool
	// sample[2] = true
	// fmt.Println(sample)

	//to add array elements while declaring array->to declare it in single line
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	//2d arrays
	nums2d := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums2d)

	/*
		in arrays:
		1. arrays are fixed in size, once we declare an array with a specific size, we cannot change its size.
		2. arrays are value types, when we assign an array to another variable, it creates a copy of the array. So, if we change the value of one array, it does not affect the other array.
		3. arrays are indexed, we can access the elements of an array using their index. The index starts from 0 and goes up to n-1, where n is the size of the array.
		4. arrays can be multi-dimensional, we can create arrays of arrays to create multi-dimensional arrays.
		5. arrays are of fixed type, we can only store elements of the same type in an array.
		6. arrays are stored in contiguous memory locations, which makes them efficient for accessing elements using their index.

	*/
}
