package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var backtrack func(level, sum int)

	res := [][]int{}
	now := []int{}
	backtrack = func(level, sum int) {
		if level == len(candidates) || sum >= target {
			if sum == target {
				res = append(res, append([]int{}, now...))
			}
			return
		}

		for i := level; i < len(candidates); i++ {
			if i > level && candidates[i] == candidates[i-1] {
				continue
			}

			now = append(now, candidates[i])
			backtrack(i+1, sum+candidates[i])
			now = now[:len(now)-1]
		}
	}

	backtrack(0, 0)

	return res
}
