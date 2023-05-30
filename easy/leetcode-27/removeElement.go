package main

import (
	"fmt"
)

func main() {
	fmt.Println("removeElement()", removeElement([]int{3, 2, 2, 3}, 3))
	// fmt.Println("removeElement()", removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))

}

func removeElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}

	// remain := []int{}
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		} else {
			nums[i] = -1
		}
	}
	fmt.Println(nums)
	return count
}
