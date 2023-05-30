package main

import (
	"fmt"
)

func main() {
	//
	fmt.Println("first Occurene in string at >> ", StrStr("sadbutsad", "sad"))
	fmt.Println("first Occurene in string at >> ", StrStr("hello", "ll"))
	fmt.Println("first Occurene in string at >> ", StrStr("leetcode", "leeto"))
	fmt.Println("first Occurene in string at >> ", StrStr("aaaaa", "bba"))

}

func StrStr(haystack string, needle string) int {

	// fmt.Println("first Occurene in string")
	// left := 0
	// right := 0

	fmt.Println(haystack, "  len(haystack) ", len(haystack))
	fmt.Println(needle, " len(needle) ", len(needle))
	fmt.Println(needle, " len(needle)+1 ", len(needle)+1)

	a1 := len(haystack)
	a2 := len(needle) + 1
	fmt.Println(a1, " ======= ", a2)
	diff := a1 - a2
	fmt.Println("len(haystack)-len(needle)+1 ", diff)

	for i := 0; i < len(haystack)-len(needle); i++ {
		// haystack[i:i+len(needle)]
		// haystack[i-i+len(needle)]
		fmt.Println(" for loop ", string(haystack[i:i+len(needle)]))
	}

	return -1
}
