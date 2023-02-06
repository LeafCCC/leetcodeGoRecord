package leetcode

import "sort"

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)

	res := 0

	for _, val := range coins {
		if val-1 > res {
			return res + 1
		}
		res += val
	}

	return res + 1
}
