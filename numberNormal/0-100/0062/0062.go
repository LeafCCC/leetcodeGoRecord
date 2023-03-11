package main

func uniquePaths(m int, n int) int {
	//f(m, n) = f(m - 1, n) + f(m, n - 1)
	if m == 1 || n == 1 {
		return 1
	}
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
		f[i][0] = 1
	}

	for j := range f[0] {
		f[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}

	return f[m-1][n-1]
}
