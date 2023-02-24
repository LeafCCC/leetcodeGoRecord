package leetcode

import "sort"

func minimumOperations(nums []int) int {
	sort.Ints(nums)
	res := 0
	i := 0
	for i < len(nums) && nums[i] == 0 {
		i++
	}

	if i == 0 {
		res++
		i++
	}

	for i < len(nums) {
		if nums[i] != nums[i-1] {
			res++
		}
		i++
	}

	return res
}
