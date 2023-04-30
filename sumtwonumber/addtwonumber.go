package sum

func List(nums []int, target int) []int {

	index := 0
	for i := 0; i < len(nums); i++ {
		index++
		for j := index; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return []int{}

}
