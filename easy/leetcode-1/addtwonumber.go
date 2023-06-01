package main

import "fmt"

func main() {

	nums1 := []int{2, 7, 11, 15}
	result1 := twoSum(nums1, 9)
	fmt.Println("result = ", result1)
	nums2 := []int{3, 2, 4}
	result2 := twoSum(nums2, 6)
	fmt.Println("result = ", result2)
	nums3 := []int{3, 2, 3}
	result3 := twoSum(nums3, 6)
	fmt.Println("result = ", result3)
	nums4 := []int{2, 5, 5, 11}
	result4 := twoSum(nums4, 10)
	fmt.Println("result = ", result4)
	nums5 := []int{-1, -2, -3, -4, -5}
	result5 := twoSum(nums5, -8)
	fmt.Println("result = ", result5)
}

func twoSum(nums []int, target int) []int {

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
