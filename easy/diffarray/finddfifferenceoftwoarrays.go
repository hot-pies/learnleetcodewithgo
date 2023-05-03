package main

import (
	"fmt"
)

func FindDifference(nums1 []int, nums2 []int) [][]int {

	fmt.Printf("Array 1 : %v Array 2 : %v\n\n", nums1, nums2)

	s1 := make(map[int]bool)
	s2 := make(map[int]bool)

	a1 := []int{}
	a2 := []int{}
	nums3 := [][]int{}

	for index := range nums1 {
		s1[nums1[index]] = true
	}

	for index := range nums2 {
		s2[nums2[index]] = true
	}

	// fmt.Println(s1)
	for i := range s1 {
		fmt.Println("s1[i] : ", i)
		// fmt.Println("nums1[i] : ", s1[i])
		// fmt.Println("nums1[i] : ", s2[nums1[i]])
		if _, ok := s2[i]; !ok {
			// fmt.Println("nums1[i] : ", nums1[i])
			a1 = append(a1, i)
		}
	}

	for i := range s2 {
		// fmt.Println("nums1[i] : ", s1[nums2[i]])
		if _, ok := s1[i]; !ok {
			// fmt.Println("nums2[i] : ", nums2[i])
			a2 = append(a2, i)
		}
	}

	// fmt.Println(a1)
	// fmt.Println(a2)
	nums3 = append(nums3, a1)
	nums3 = append(nums3, a2)
	fmt.Println(nums3)
	return nums3
}
