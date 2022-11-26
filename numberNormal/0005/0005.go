package leetcode

//中心扩散算法
func longestPalindrome(s string) string {
	res := ""
	n := len(s)
	for i := 0; i < 2*n-1; i++ {
		start, end := i/2, (i+1)/2
		for start >= 0 && end < n {
			if s[start] == s[end] {
				start--
				end++
			} else {
				break
			}
		}
		if end-start-1 > len(res) {
			res = s[start+1 : end]
		}
	}

	return res
}
