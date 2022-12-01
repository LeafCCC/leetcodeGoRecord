package leetcode

func match(a string, b string) bool {
	for i := range b {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if match(haystack[i:], needle) {
			return i
		}
	}

	return -1
}
