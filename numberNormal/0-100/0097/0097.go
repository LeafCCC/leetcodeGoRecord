package leetcode

func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)

	if n1+n2 != n3 {
		return false
	}

	f := make([][]bool, n1+1)
	for i := range f {
		f[i] = make([]bool, n2+1)
	}

	f[0][0] = true

	for j := 1; j <= n2; j++ {
		if s2[j-1] == s3[j-1] {
			f[0][j] = f[0][j-1]
		}
	}

	for i := 1; i <= n1; i++ {
		if s1[i-1] == s3[i-1] {
			f[i][0] = f[i-1][0]
		}
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if (s1[i-1] == s3[i+j-1] && f[i-1][j]) || (s2[j-1] == s3[i+j-1] && f[i][j-1]) {
				f[i][j] = true
			}
		}
	}

	return f[n1][n2]

}
