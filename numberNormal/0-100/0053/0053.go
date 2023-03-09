package leetcode

func maxSubArray(nums []int) int {
	n := len(nums)

	dp := make([]int, n+1)

	res := nums[0]

	for i := range nums {
		dp[i+1] = max(dp[i]+nums[i], nums[i])

		res = max(res, dp[i+1])
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
