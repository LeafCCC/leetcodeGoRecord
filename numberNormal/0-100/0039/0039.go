package leetcode

import "sort"

func combinationSum(candidates []int, target int) (res [][]int) {
	sort.Ints(candidates)

	now := []int{}
	var backTrack func(level int, sum int)

	backTrack = func(level int, sum int) {
		if sum > target {
			return
		} else if sum == target {
			res = append(res, append([]int{}, now...))
			return
		}

		for i := level; i < len(candidates); i++ {
			now = append(now, candidates[i])
			backTrack(i, sum+candidates[i])
			now = now[:len(now)-1]
		}
	}

	backTrack(0, 0)
	return
}
