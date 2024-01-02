package main

import (
	"fmt"
)

type stack []string

func (s *stack) push(e string) {
	*s = append(*s, e)
	// fmt.Println(s)
}

func (s *stack) pop() string {
	if s.isEmpty() {
		return ""
	}

	top := (*s)[len(*s)-1]
	// fmt.Println(top)
	*s = (*s)[:len(*s)-1]
	// fmt.Println(s)
	return top
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func IsValid(s string) bool {

	// fmt.Println(" array = ", s)
	if len(s) < 1 {
		return false
	}

	var stack = &stack{}
	isValid := false
	chars := []rune(s)
	for i := range chars {
		// fmt.Println(" Char = ",chars[i])
		if chars[i]==40 || chars[i]==91 || chars[i]==123{
			// fmt.Println("PUSH")
			stack.push(string(chars[i]))	
			fmt.Println(*stack)
		}else{
			fmt.Println("ELSE")
			if !stack.isEmpty(){
				s:=[]rune(stack.pop())
				switch 	chars[i]{
				case 41:
					if chars[i]-s[0]==1{
						isValid=true
						fmt.Println(" isValid : ",isValid)
					}else{
						isValid = false
						fmt.Println("41 Else isValid : ",isValid)
						return isValid
					}
					fmt.Println("41 = ",string(chars[i]), "pop : ",s[0])
				case 93:
					if chars[i]-s[0]==2{
						isValid=true
						fmt.Println(" isValid : ",isValid)
					}else{
						isValid = false
						fmt.Println("93 Else isValid : ",isValid)
						return isValid
					}
					fmt.Println("93 = ",string(chars[i]), "pop : ",s[0])
				case 125:
					if chars[i]-s[0]==2{
						isValid=true
						fmt.Println(" isValid : ",isValid)
					}else{
						isValid = false
						fmt.Println("125 Else isValid : ",isValid)
						return isValid
					}
					fmt.Println("125 = ",string(chars[i]), "pop : ",s[0])
				default:
					return false
				}
			fmt.Println("POST",*stack)
			}else{
				return false
			}
		}
	}
	if !stack.isEmpty(){
		return false
	}
	return isValid
}
