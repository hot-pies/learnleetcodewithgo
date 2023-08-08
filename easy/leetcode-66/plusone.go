package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Print(PlusOne([]int{1, 2, 9, 9}))

	fmt.Print(PlusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}))
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
