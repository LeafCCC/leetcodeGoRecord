package leetcode

import "sort"

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	num := append(nums1, nums2...)

	sort.Slice(num, func(i, j int) bool {
		return num[i][0] < num[j][0]
	})

	i := 1

	for i < len(num) {
		if num[i][0] == num[i-1][0] {
			num[i][1] += num[i-1][1]
			num = append(num[:i-1], num[i:]...)
		} else {
			i++
		}
	}

	return num
}
