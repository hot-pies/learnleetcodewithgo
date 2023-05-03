package longestpalindromicsubstring

import "fmt"

func LongestPalindrome(s string) string {

	s1 := []rune(s)

	// newpalindroms := []string{}
	// babad
	newpalindroms := ""
	// right := ""
	// lenght := 0
	r := 0
	// l := 0
	index := len(s1) - 1

	for index >= 0 {
		// ("cbbd")
		// fmt.Println(string(s1[r]), " == ", string(s1[index]))
		if string(s1[r]) == string(s1[index]) {
			newpalindroms += string(s1[r])
			// right += string(s1[index])
			fmt.Println("Inside if : ", r, index)
			// r++
			index--
		} else {
			index--
			r++
		}
		fmt.Println(r, " == ", index)
	}
	return newpalindroms
}
