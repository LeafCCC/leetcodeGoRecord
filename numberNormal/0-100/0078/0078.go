package leetcode

func subsets(nums []int) [][]int {
	res := [][]int{[]int{}}
	tmp := []int{}

	var backTrack func(int)

	backTrack = func(now int) {
		for i := now; i < len(nums); i++ {
			tmp = append(tmp, nums[i])
			res = append(res, append([]int{}, tmp...))
			backTrack(i + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}

	backTrack(0)

	return res
}
