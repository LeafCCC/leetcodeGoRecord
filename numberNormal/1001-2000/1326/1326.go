package leetcode

import "sort"

func minTaps(n int, ranges []int) int {
	scopes := [][]int{}
	for i := range ranges {
		left, right := i-ranges[i], i+ranges[i]
		if left == right {
			continue
		}
		scopes = append(scopes, []int{left, right})
	}
	sort.Slice(scopes, func(i, j int) bool {
		return scopes[i][0] < scopes[j][0]
	})

	now := 0
	si := 0

	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	res := 0
	for si < len(scopes) && now < n {
		if scopes[si][0] > now {
			return -1
		}

		res++
		maxR := now
		for si < len(scopes) && scopes[si][0] <= now {
			maxR = max(maxR, scopes[si][1])
			si++
		}
		now = maxR
	}

	if now < n {
		return -1
	}
	return res

}
