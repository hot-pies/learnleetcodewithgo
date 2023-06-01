package main

import (
	"fmt"
)

func main() {
	PlusOne([]int{1, 2, 3, 9})
}

func PlusOne(arr []int) []int {
	fmt.Println("Plus One ", arr)

	plus := 1
	if len(arr) == 1 {
		return []int{arr[0] + plus}
	}

	if 10-arr[len(arr)-1] == 1 {
		fmt.Println(arr[len(arr)-1])
	}

	return []int{}
}
