package leetcode

// x, y 是蛇尾位置 s表示状态 0平躺 1竖直(x, y, s)
type state struct{ x, y, s int }

// 分别代表右移 下移 转向
var move = []state{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

func minimumMoves(grid [][]int) int {
	q := []state{{0, 0, 0}}

	n := len(grid)
	vis := make([][][2]bool, n)
	for i := range vis {
		vis[i] = make([][2]bool, n)
	}
	vis[0][0][0] = true

	for step := 1; len(q) > 0; step++ {
		tmp := q
		q = nil
		for _, t := range tmp {
			for _, m := range move {
				x1, y1, s := t.x+m.x, t.y+m.y, t.s^m.s
				x2, y2 := x1+s, y1+(s^1)

				if x2 >= n || y2 >= n || vis[x1][y1][s] {
					continue
				}

				if grid[x1][y1] == 0 && grid[x2][y2] == 0 && (m.s == 0 || grid[x1+1][y1+1] == 0) {
					if x1 == n-1 && y1 == n-2 && s == 0 {
						return step
					}
					q = append(q, state{x1, y1, s})
					vis[x1][y1][s] = true
				}
			}
		}
	}

	return -1
}
