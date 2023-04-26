package leetcode

func isPalindrome(s string) bool {
	a := []byte(s)
	b := []byte{}

	for i := range a {
		switch {
		case a[i] >= 'a' && a[i] <= 'z':
			b = append(b, a[i])
		case a[i] >= '0' && a[i] <= '9':
			b = append(b, a[i])
		case a[i] >= 'A' && a[i] <= 'Z':
			b = append(b, byte(a[i]+'a'-'A'))
		}
	}

	c := make([]byte, len(b))
	for i := range b {
		c[len(b)-1-i] = b[i]
	}

	return string(b) == string(c)
}
