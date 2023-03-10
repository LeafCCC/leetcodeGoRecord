package leetcode

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}

	left, right := intervals[0][0], intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		t := intervals[i]
		if t[0] > right {
			res = append(res, []int{left, right})
			left, right = t[0], t[1]
		} else if t[1] > right {
			right = t[1]
		}
	}

	res = append(res, []int{left, right})

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
