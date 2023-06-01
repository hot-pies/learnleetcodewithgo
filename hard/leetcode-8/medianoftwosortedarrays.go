package medianofarray

import (
	"fmt"
	"sort"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	median := 0.0
	if !sort.IntsAreSorted(nums1) && sort.IntsAreSorted(nums2) {
		return median
	}

	nums3 := append(nums1, nums2...)
	sort.Ints(nums3)
	fmt.Println(nums3)
	if len(nums3)%2 != 0 {
		return float64(nums3[len(nums3)/2])
	} else {
		fmt.Println("Even length ", nums3[(len(nums3)/2)-1])
		mid1 := float64(nums3[(len(nums3)/2)-1])
		mid2 := float64(nums3[len(nums3)/2])
		fmt.Printf("mid1 : %v  mid2 : %v", mid1, mid2)

		fmt.Println()
		fmt.Println(mid1 + mid2)
		if mid1 > 0 && mid2 > 0 {
			return float64((mid1 + mid2) / 2)
		} else {
			return median
		}

	}
}
