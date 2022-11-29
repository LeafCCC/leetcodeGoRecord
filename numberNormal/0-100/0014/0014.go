package leetcode

func longestCommonPrefix(strs []string) string {
	t := []byte(strs[0])

	for i := 1; i < len(strs); i++ {
		now := []byte(strs[i])

		for j := range t {
			if j > len(now)-1 {
				t = now
				break
			}
			if t[j] != now[j] {
				t = t[:j]
				break
			}
		}

		if len(t) == 0 {
			return ""
		}
	}

	return string(t)
}
