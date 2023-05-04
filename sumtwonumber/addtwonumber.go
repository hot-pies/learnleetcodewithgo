package main

import "fmt"

func List(nums []int, target int) []int {

	right := len(nums) - 1
	left := 0
	// fmt.Printf("Left: %v Right :%v", left, right)
	for left < right {
		// fmt.Printf("Left: %v Right :%v", left, right)
		sum := nums[left] + nums[right]
		fmt.Println("sum : ", sum, "target : ", target)
		if sum == target {
			return []int{left, right}
		} else if nums[right] > target {
			right--
		} else if nums[left] < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}

	return nil

	// index := 0
	// for i := 0; i < len(nums); i++ {
	// 	index++
	// 	for j := index; j < len(nums); j++ {
	// 		if (nums[i] + nums[j]) == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }
	// return []int{}
}
