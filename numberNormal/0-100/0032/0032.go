package leetcode

//动态规划
func longestValidParentheses(s string) (res int) {
	n := len(s)
	if n < 2 {
		return
	}

	//注意dp[i]的意思是以s[i-1]结尾的最长子串的长度
	dp := make([]int, n+1)

	for i := 1; i < n; i++ {
		if s[i-1] == '(' && s[i] == ')' {
			dp[i+1] = dp[i-2+1] + 2
		} else if s[i-1] == ')' && s[i] == ')' {
			if i-1-dp[i-1+1] >= 0 && s[i-1-dp[i-1+1]] == '(' {
				dp[i+1] = dp[i-1+1] + 2 + dp[i-2-dp[i-1+1]+1]
			}
		}

		if dp[i+1] > res {
			res = dp[i+1]
		}
	}

	return
}
