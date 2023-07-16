package leetcode

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	rcd := make(map[string]bool)
	for _, v := range wordDict {
		rcd[v] = true
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] = dp[j] && rcd[s[j:i]]
			if dp[i] {
				break
			}
		}
	}

	return dp[n]
}
