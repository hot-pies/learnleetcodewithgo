package main

import (
	"fmt"
)

func main() {

	nums1 := []int{2, 7, 11, 15}
	result1 := List(nums1, 9)
	fmt.Println("result = ", result1)
	nums2 := []int{3, 2, 4}
	result2 := List(nums2, 6)
	fmt.Println("result = ", result2)
	nums3 := []int{3, 2, 3}
	result3 := List(nums3, 6)
	fmt.Println("result = ", result3)
	nums4 := []int{2, 5, 5, 11}
	result4 := List(nums4, 10)
	fmt.Println("result = ", result4)
	nums5 := []int{-1, -2, -3, -4, -5}
	result5 := List(nums5, -8)
	fmt.Println("result = ", result5)
}
