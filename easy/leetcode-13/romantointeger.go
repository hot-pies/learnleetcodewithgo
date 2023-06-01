package main

import "fmt"

func main() {
	// fmt.Println(RomanToInt("III"))
	// fmt.Println(RomanToInt("LVIII"))
	fmt.Println(RomanToInt("MCMXCIV"))

}

func RomanToInt(s string) int {

	roman := make(map[string]int)
	roman["I"] = 1
	roman["V"] = 5
	roman["X"] = 10
	roman["L"] = 50
	roman["C"] = 100
	roman["D"] = 500
	roman["M"] = 1000

	chars := []rune(s)

	sum := 0
	for i := range chars {
		// fmt.Printf("%v => %v ", string(chars[i]), roman[string(chars[i])])
		if i > 0 && roman[string(chars[i])] > roman[string(chars[i-1])] {
			// fmt.Println("==if== ", roman[string(chars[i])]-roman[string(chars[i-1])])
			sum = sum + (roman[string(chars[i])] - (roman[string(chars[i-1])] + roman[string(chars[i-1])]))
			// fmt.Println("Sum = ", sum)
		} else {
			// fmt.Printf("else == %v => %v ", string(chars[i]), roman[string(chars[i])])
			sum = sum + roman[string(chars[i])]
			// fmt.Println("Sum == ", sum)
		}
	}

	return sum
}
