package leetcode

func largestMerge(word1 string, word2 string) string {
	var f func(string, string, string) string
	f = func(a string, b string, now string) string {
		if len(a) == 0 && len(b) == 0 {
			return now
		} else if len(a) == 0 {
			return now + b
		} else if len(b) == 0 {
			return now + a
		}

		if a > b {
			return f(a[1:], b, now+a[0:1])
		}
		return f(a, b[1:], now+b[0:1])

	}

	return f(word1, word2, "")
}
