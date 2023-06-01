package longestsustring

func LengthOfLongestSubstring(s string) int {

	if len(s) == 1 {
		return len(s)
	}
	var charr = []rune(s)
	var subset = make(map[string]bool)
	// Input "pwwkew"
	var r = 0
	var l = 0
	var length = 0
	index := 0
	// fmt.Println("string(charr[r]) : ", string(charr[r]))
	for len(charr) > index {
		// fmt.Println("string(charr[r]) : ", string(charr[r]))
		// fmt.Printf("Right : %v Left : %v new value : %v ", r, l, string(charr[r]))
		// fmt.Println()
		// fmt.Printf("Right : %v Left : %v Index : %v ", r, l, index)

		if _, ok := subset[string(charr[r])]; !ok {
			// fmt.Println(" Key No exist Add : ", string(charr[r]))
			sublen := (r - l) + 1
			// fmt.Println(" sublen : ", sublen)

			if length < sublen {
				length = sublen
			}
			// fmt.Println(" length : ", length)
			// subset[string(charr[index])] = string(charr[index])
			subset[string(charr[r])] = true
			// fmt.Println("After + add  substr : ", subset)
			// fmt.Println("subset := ", subset)
			r++
			index++
		} else {
			// fmt.Println("inside else")
			if _, ok := subset[string(charr[r])]; ok {
				// fmt.Println(" Delete : ", subset[string(charr[l])])
				delete(subset, string(charr[l]))
				// fmt.Println("- After delete  substr : ", subset)
				l++
			}
		}
		// fmt.Println("Value of I : ", index, " substring : ", subset)

	}
	return length
	// qrsvbspk
}

// func LengthOfLongestSubstring(s string) int {

// 	// fmt.Print("input", strings.Split(s, ","), " ")
// 	char := []rune(s)
// 	var sb1 = ""
// 	var sb2 = ""
// 	for i := 0; i < len(char); i++ {
// 		if !strings.Contains(sb1, string(s[i])) {
// 			// fmt.Println(string(s[i]))
// 			sb1 = sb1 + string(s[i])
// 		} else {
// 			// fmt.Println("else : ", sb)
// 			sb2 = sb1
// 			sb1 = ""
// 		}
// 		// fmt.Print("  Existing : ", sb)
// 		// fmt.Print("  old sub : ", oldsb)
// 	}
// 	if len(sb2) > len(sb1) {
// 		fmt.Println("return")
// 		return len(sb2)
// 	}
// 	fmt.Print(sb1, " ")
// 	return len(sb1) - 1
// }
