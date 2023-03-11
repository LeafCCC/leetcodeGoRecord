package main

func uniquePathsWithObstacles(grid [][]int) int {
	//f(m, n) = f(m-1, n) + f(m, n-1)
	if grid[0][0] == 1 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}

	f[0][0] = 1
	for j := 1; j < n; j++ {
		if grid[0][j] == 1 {
			f[0][j] = 0
		} else {
			f[0][j] = f[0][j-1]
		}
	}

	for i := 1; i < m; i++ {
		if grid[i][0] == 1 {
			f[i][0] = 0
		} else {
			f[i][0] = f[i-1][0]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == 1 {
				f[i][j] = 0
			} else {
				f[i][j] = f[i-1][j] + f[i][j-1]
			}
		}
	}

	return f[m-1][n-1]
}
