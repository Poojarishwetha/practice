package main

import (
	"fmt"
)

//slices - a slice is a dynamically sized array. It is a reference type and it is built on top of arrays. A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and the capacity of the segment. The length of the segment is the number of elements in the segment, and the capacity of the segment is the number of elements in the underlying array starting from the first element in the segment. A slice can be created using the make function or by slicing an existing array or slice.
//most used concept in go

func main() {
	// //to declare slices
	// var nums []int //uninitialized slice is "nil"
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))
	// var vals = make([]int, 2)
	// fmt.Println(cap(vals))
	// //to add elements
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	//************************************************************
	// var nums = make([]int, 2, 5)
	// 	fmt.Println(nums)
	// 	fmt.Println(len(nums))
	// 	fmt.Println(cap(nums))
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 2)
	// nums = append(nums, 2)
	// nums = append(nums, 2)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	//****************************************************

	// nums1 := []int{}
	// // nums1 = append(nums1, 1, 1, 2, 3, 4, 5)
	// nums1 = append(nums1, 7)
	// nums1 = append(nums1, 9)
	// fmt.Println(nums1)
	// fmt.Println(len(nums1))
	// fmt.Println(cap(nums1))
	//*****************************************************************

	//slice operator
	// var nums2 = []int{1, 2, 3}
	// fmt.Println(nums2[0:1])
	// fmt.Println(nums2[:1])
	// fmt.Println(nums2[1:])
	//*******************************************************

	//slices package-> provides functions to manipulate slices. It is a part of the standard library. It provides functions to compare, copy, delete, insert, reverse, sort, and unique slices.
	// var num3 = []int{1, 2, 4, 5}
	// var num4 = []int{1, 2, 6, 5}
	// fmt.Println(slices.Equal(num3, num4))
	//******************************************************

	var num5 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(num5)
}
