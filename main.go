package main

import (
	"fmt"
	"golearning/leetcode/longestsustring"
)

func main() {

	// longestsustring.LengthOfLongestSubstring("abcabcbb")
	// longestsustring.LengthOfLongestSubstring("bbbbb")
	fmt.Println(longestsustring.LengthOfLongestSubstring("abcabcbb"))
	fmt.Println(longestsustring.LengthOfLongestSubstring("bbbbb"))
	fmt.Println(longestsustring.LengthOfLongestSubstring("pwwkew"))
	fmt.Println(longestsustring.LengthOfLongestSubstring("aab"))
	fmt.Println(longestsustring.LengthOfLongestSubstring(" "))
	fmt.Println(longestsustring.LengthOfLongestSubstring("au"))
	fmt.Println(longestsustring.LengthOfLongestSubstring("bb"))
	fmt.Println(longestsustring.LengthOfLongestSubstring("qrsvbspk"))

}

// Call Add Two Number
// func main() {

// 	// nums := []int{2, 7, 11, 15}
// 	// nums := []int{3, 2, 4}
// 	// nums := []int{3, 2, 3}
// 	nums := []int{2, 5, 5, 11}
// 	result := sum.List(nums, 10)
// 	fmt.Println("result = ", result)
// }
