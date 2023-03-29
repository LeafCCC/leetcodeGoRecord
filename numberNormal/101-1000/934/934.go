package leetcode

func shortestBridge(grid [][]int) int {
	x, y := 0, 0
	m, n := len(grid), len(grid[0])
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				x, y = i, j
				break
			}
		}
	}

	type S struct{ x, y int }

	q := []S{S{x, y}}
	island := []S{}
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	vis[x][y] = true

	dirs := []int{-1, 0, 1, 0, -1}
	for len(q) != 0 {
		tmp := []S{}
		for _, s := range q {
			x, y := s.x, s.y
			grid[x][y] = -1
			island = append(island, s)
			for i := 0; i < 4; i++ {
				new_x, new_y := x+dirs[i], y+dirs[i+1]
				if new_x >= 0 && new_x < m && new_y >= 0 && new_y < n && grid[new_x][new_y] == 1 && !vis[new_x][new_y] {
					tmp = append(tmp, S{new_x, new_y})
					vis[new_x][new_y] = true
				}
			}
		}
		q = tmp
	}

	res := -2

	for len(island) != 0 {
		res++
		tmp := []S{}
		for _, s := range island {
			x, y := s.x, s.y
			if grid[x][y] == 1 {
				return res
			}
			vis[x][y] = true
			for i := 0; i < 4; i++ {
				new_x, new_y := x+dirs[i], y+dirs[i+1]
				if new_x >= 0 && new_x < m && new_y >= 0 && new_y < n && !vis[new_x][new_y] {
					tmp = append(tmp, S{new_x, new_y})
				}
			}
		}
		island = tmp
	}

	return res
}
