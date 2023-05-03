package leetcode

func partition(s string) [][]string {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = f[i+1][j-1] && s[i] == s[j]
		}
	}

	dp := make([][][]string, len(s))
	dp[0] = [][]string{[]string{s[0:1]}}
	for i := 1; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if f[j][i] {
				if j == 0 {
					dp[i] = append(dp[i], []string{s[j : i+1]})
				} else {
					for t := range dp[j-1] {
						tmp := append([]string{}, dp[j-1][t]...)
						tmp = append(tmp, s[j:i+1])
						dp[i] = append(dp[i], tmp)
					}
				}
			}
		}
	}
	return dp[len(s)-1]
}
