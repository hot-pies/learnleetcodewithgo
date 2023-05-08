package main

import (
	"fmt"
	"strings"
)

func StrStr(haystack string, needle string) int {
	fmt.Println("first Occurene in string")

	index := strings.Index(haystack, needle)

	fmt.Println(index)

	return index
}
