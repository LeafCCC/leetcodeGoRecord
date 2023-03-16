package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{[]int{}}
	tmp := []int{}
	sort.Ints(nums)

	var backTrack func(int)

	backTrack = func(now int) {
		for i := now; i < len(nums); i++ {
			if i > now && nums[i] == nums[i-1] {
				continue
			}
			tmp = append(tmp, nums[i])
			res = append(res, append([]int{}, tmp...))
			backTrack(i + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}

	backTrack(0)
	return res
}
