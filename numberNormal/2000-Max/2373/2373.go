package leetcode

func largestLocal(grid [][]int) [][]int {
	n := len(grid)

	res := make([][]int, n-2)
	for i := range res {
		res[i] = make([]int, n-2)
	}

	dirs := []int{0, -1, 0, 1, 0, 1, 1, -1, -1, 1}

	for i := range res {
		for j := range res[i] {
			tmp := grid[i+1][j+1]
			for d := range dirs[:len(dirs)-1] {
				tmp = max(tmp, grid[i+1+dirs[d]][j+1+dirs[d+1]])
			}
			res[i][j] = tmp
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
