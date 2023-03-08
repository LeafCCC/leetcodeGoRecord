package leetcode

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	res := 0

	right := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= right {
			right = intervals[i][1]
		} else {
			res++
		}
	}

	return res
}
