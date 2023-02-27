package leetcode

func movesToMakeZigzag(nums []int) int {
	res1, res2, n := 0, 0, len(nums)

	//偶数索引的情况
	for i := 1; i < n-1; i += 2 {
		res1 += cal(nums[i], min(nums[i-1], nums[i+1]))
	}

	//奇数索引的情况
	for i := 2; i < n-1; i += 2 {
		res2 += cal(nums[i], min(nums[i-1], nums[i+1]))
	}

	if n > 1 {
		res2 += cal(nums[0], nums[1])

		if n%2 == 0 {
			res1 += cal(nums[n-1], nums[n-2])
		} else {
			res2 += cal(nums[n-1], nums[n-2])
		}
	}

	return min(res1, res2)
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
func cal(a, b int) int {
	if a-b >= 0 {
		return a - b + 1
	}
	return 0
}
