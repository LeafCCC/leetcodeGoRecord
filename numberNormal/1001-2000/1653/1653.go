package leetcode

func minimumDeletions(s string) int {
	dp, bcount := 0, 0

	for _, c := range s {
		if c == 'a' {
			dp = min(dp+1, bcount)
		} else {
			bcount++
		}
	}

	return dp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
