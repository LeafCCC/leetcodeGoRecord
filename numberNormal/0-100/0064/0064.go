package main

func minPathSum(grid [][]int) int {
	//f[i][j] = grid[i][j] + min(f[i-1][j], f[i][j-1])

	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}

	f[0][0] = grid[0][0]
	for j := 1; j < n; j++ {
		f[0][j] = f[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		f[i][0] = f[i-1][0] + grid[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[i][j] = grid[i][j] + min(f[i-1][j], f[i][j-1])
		}
	}

	return f[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
