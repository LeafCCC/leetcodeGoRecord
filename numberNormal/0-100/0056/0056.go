package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		t := intervals[i]

		if t[0] > right {
			res = append(res, []int{left, right})
			left = t[0]
			right = t[1]
		} else if t[1] > right {
			right = t[1]
		}
	}
	res = append(res, []int{left, right})
	return res
}
