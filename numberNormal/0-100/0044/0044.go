<<<<<<< HEAD
package leetcode

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for j := range p {
		if p[j] == '*' {
			dp[0][j+1] = dp[0][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[j] == '?' {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				dp[i+1][j+1] = dp[i+1][j] || dp[i][j+1]
			} else {
				dp[i+1][j+1] = dp[i][j] && s[i] == p[j]
			}
		}
	}

	return dp[m][n]
}
=======
package leetcode

func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    dp := make([][]bool, m + 1)
    for i := range dp{
        dp[i] = make([]bool, n + 1)
    }
    
    dp[0][0] = true
    for j := range p{
        if p[j] == '*'{
            dp[0][j+1] = dp[0][j]
        }
    }
    
    for i := 0; i < m; i++{
        for j := 0; j < n; j++{
            if p[j] == '?'{
                dp[i+1][j+1] = dp[i][j]
            }else if p[j] == '*'{
                dp[i+1][j+1] = dp[i+1][j] || dp[i][j+1]
            }else{
                dp[i+1][j+1] = dp[i][j] && s[i] == p[j]
            }
        }
    }
    
    return dp[m][n]
}
>>>>>>> 89c72de768249774a3ee20e7263c2916d0b8f194
