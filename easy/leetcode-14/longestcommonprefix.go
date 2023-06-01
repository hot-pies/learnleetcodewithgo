package main

import (
	"fmt"
)

func main() {
	// fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	// fmt.Println(LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	// fmt.Println(LongestCommonPrefix([]string{""}))
	// fmt.Println(LongestCommonPrefix([]string{"a"}))
	// fmt.Println(LongestCommonPrefix([]string{"ab", "a"}))
	fmt.Println(LongestCommonPrefix([]string{"flower", "flower", "flower", "flower"}))

}

func LongestCommonPrefix(strs []string) string {

	fmt.Println(strs)
	if len(strs) <= 0 {
		return ""
	}

	fmt.Println("arrya is not null", strs)
	prefix := ""
	// count := 0
	ch1 := ""
	for i := 0; i < len(strs[0]); i++ {
		fmt.Println("  continue with new i ", i)
		// if len(strs[0]) > 0 {
		fmt.Println("exit ", i)
		ch1 = string(strs[0][i])
		// }

		// fmt.Println("  char = ", ch1)
		for j := 1; j < len(strs); j++ {
			fmt.Println("inside second loop ", j, len(strs[j]))
			// fmt.Println("   J = ", strs[j])
			// fmt.Println("   I = ", strs[i])
			// fmt.Printf("  %v %v \n", i, j)
			if i >= len(strs[j]) || ch1 != string(strs[j][i]) {
				fmt.Println("Exit")
				return prefix
			}
		}
		prefix = prefix + ch1
	}
	fmt.Println("  for done ")
	return prefix
}

// func LongestCommonPrefix(strs []string) string {

// 	fmt.Println(strs)

// 	if len(strs) <= 0 {
// 		return ""
// 	}
// 	match := ""
// 	left := 0

// 	if len(strs) > 1 {
// 		chars := []rune(strs[0])
// 		prefix := ""
// 		if len(chars) > 1 {
// 			prefix = string(chars[0])
// 		}

// 		for i := range strs {
// 			chars1 := []rune(strs[i])
// 			pre := ""
// 			if len(chars1) > i {
// 				pre = string(chars1[i])
// 			}

// 			if prefix == pre {
// 				match = match + prefix
// 				left++
// 			}
// 			if len(chars1) > left {
// 				prefix = string(chars1[left])
// 			}
// 		}
// 	} else {
// 		return strs[0]
// 	}

// 	if left > 0 {
// 		return match
// 	} else {
// 		return ""
// 	}

// }
