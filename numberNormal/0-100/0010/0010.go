package leetcode

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)

	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	//(1) p[j] != '*'
	//    s[i] == p[j] || p[j] == '.', dp[i][j] = dp[i- 1][j - 1]
	//    s[i] != p[j], dp[i][j] = false
	//(2) p[j] == '*'
	//    s[i] == p[j-1] || p[j-1] == '.'
	//    dp[i][j] = dp[i-1][j] || dp[i][j-2]
	//    else dp[i][j] = dp[i][j-2]

	dp[0][0] = true
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[j] != '*' {
				if s[i] == p[j] || p[j] == '.' {
					dp[i+1][j+1] = dp[i+1-1][j+1-1]
				} else {
					dp[i+1][j+1] = false
				}
			} else {
				if s[i] == p[j-1] || p[j-1] == '.' {
					//匹配一个 或者一个都不匹配
					dp[i+1][j+1] = dp[i+1-1][j+1] || dp[i+1][j+1-2]
				} else {
					dp[i+1][j+1] = dp[i+1][j+1-2]
				}
			}
		}
	}

	return dp[m][n]
}
