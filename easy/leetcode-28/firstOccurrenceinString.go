package main

import (
	"fmt"
	"strings"
)

func StrStr(haystack string, needle string) int {
	fmt.Println("first Occurene in string")

	// match := -1
	// index := 0
	// left := 0
	result := -1
	index := strings.Index(haystack, needle)

	fmt.Println(index)

	return result
}
