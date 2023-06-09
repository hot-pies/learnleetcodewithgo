package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(PlusOne([]int{1, 2, 9, 9}))
}

func PlusOne(arr []int) []int {

	fmt.Println(arr)

	number := ""

	for i := 0; i < len(arr); i++ {
		temp := fmt.Sprint(arr[i])
		number = number + temp
	}

	temp, _ := strconv.Atoi(number)

	temp = temp + 1

	number = fmt.Sprint(temp)

	// fmt.Println(number)

	digit := []int{}

	for i := range number {
		temp, _ := strconv.Atoi(string(number[i]))
		// fmt.Println(temp)
		digit = append(digit, temp)
	}

	// fmt.Print(digit)
	return digit
}
