package leetcode

func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle))
	dp[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = min(dp[j], dp[j-1]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}
	res := dp[0]
	for _, val := range dp {
		res = min(val, res)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
