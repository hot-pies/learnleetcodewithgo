package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {

	str := fmt.Sprint(x)
	val := []rune(str)

	mid := len(val) / 2
	right := len(val) - 1
	left := 0

	for i := 0; i < mid; i++ {
		if left <= mid && right >= mid && val[left] == val[right] {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}
