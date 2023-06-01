package main

import (
	"fmt"
	"golearning/learnleetcode/datastructure"
)

type stk []string

func (s *stk) push(element string) {
	*s = append(*s, element)
}

func (s *stk) pop() string {
	if len(*s) == 0 {
		return ""
	}

	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func Comp(str1 string, str2 string) bool {
	fmt.Println(str2, str2)
	var stack = &datastructure.Stack{}

	stack.Push("world")
	stack.Push("hello ")
	fmt.Printf(" %v", stack.Pop())
	fmt.Printf(" %v\n", stack.Pop())

	return str2 == str2
}
