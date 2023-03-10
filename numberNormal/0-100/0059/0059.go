package leetcode

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	x, y := 0, 0
	dir := []int{0, 1, 0, -1, 0}
	nowDir := 0

	for i := 1; i <= n*n; i++ {
		res[x][y] = i

		tx := x + dir[nowDir%4]
		ty := y + dir[(nowDir+1)%4]

		if tx < 0 || tx >= n || ty < 0 || ty >= n || res[tx][ty] != 0 {
			nowDir++
			tx = x + dir[nowDir%4]
			ty = y + dir[(nowDir+1)%4]
		}
		x, y = tx, ty
	}

	return res
}
