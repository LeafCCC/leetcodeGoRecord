package leetcode

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 1; i < n; i++ {
		a := int(s[i-1]) - int('0')
		b := int(s[i]) - int('0')

		if b == 0 {
			if a == 1 || a == 2 {
				dp[i+1] = dp[i-1]
			} else {
				return 0
			}
		} else {
			switch a {
			case 1:
				dp[i+1] = dp[i] + dp[i-1]
			case 2:
				if b < 7 {
					dp[i+1] = dp[i] + dp[i-1]
				} else {
					dp[i+1] = dp[i]
				}
			default:
				dp[i+1] = dp[i]
			}
		}
	}

	return dp[n]

}
