package leetcode

func permute(nums []int) [][]int {
	res := [][]int{}

	var backTracking func(step int)

	backTracking = func(step int) {
		if step == len(nums)-1 {
			res = append(res, append([]int{}, nums...))
			return
		}

		for i := step; i < len(nums); i++ {
			nums[step], nums[i] = nums[i], nums[step]
			backTracking(step + 1)
			nums[step], nums[i] = nums[i], nums[step]
		}
	}

	backTracking(0)
	return res
}
