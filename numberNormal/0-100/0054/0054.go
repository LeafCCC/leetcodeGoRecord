package leetcode

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	dir := []int{0, 1, 0, -1, 0}
	now := 0 //现在的方向
	res := []int{}
	x, y := 0, 0

	for {
		res = append(res, matrix[x][y])
		vis[x][y] = true

		x2, y2 := x+dir[now%4], y+dir[(now%4)+1]

		if x2 >= m || y2 >= n || x2 < 0 || y2 < 0 || vis[x2][y2] {
			now++
			x2, y2 = x+dir[now%4], y+dir[(now%4)+1]
			if x2 >= m || y2 >= n || x2 < 0 || y2 < 0 || vis[x2][y2] {
				break
			}
		}

		x, y = x2, y2
	}

	return res
}
