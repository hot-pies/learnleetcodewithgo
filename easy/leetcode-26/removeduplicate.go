package main

import (
	"fmt"
)

func RemoveDuplicates(nums []int) int {
	fmt.Println("Replace Duplicate from array", nums)

	if len(nums) <= 0 {
		return 0
	}

	index := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}

	// length := len(nums) - 1
	// for length > left {
	// 	// fmt.Printf("length %v right %v element %v\n", length, right, nums[index])
	// 	// 1, 2, 3, 2, 1, 3, 1
	// 	if index == nums[length] {
	// 		nums[length] = 101
	// 	}

	// 	if index == nums[right] {
	// 		nums[right] = 101
	// 	}
	// 	if right < length {
	// 		right++
	// 	}else{

	// 	}

	// 	// fmt.Printf("length %v right %v\n", length, right)
	// 	if right == length {
	// 		left++
	// 		index = nums[left]
	// 		right = left + 1
	// 	}
	// }

	// nums = sort.IntSlice(nums)
	// fmt.Println("left = ", left)

	return index

}
