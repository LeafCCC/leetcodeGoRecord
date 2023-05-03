package leetcode

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	dirs := []int{-1, 0, 1, 0, -1}

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' || vis[i][j] {
				continue
			}

			tag := false
			list := [][]int{[]int{i, j}}
			now := [][]int{[]int{i, j}}
			vis[i][j] = true

			for len(now) != 0 {
				tmp := [][]int{}
				for _, v := range now {
					x, y := v[0], v[1]
					for l := 0; l < 4; l++ {
						nx, ny := x+dirs[l], y+dirs[l+1]
						if nx < 0 || nx >= m || ny < 0 || ny >= n {
							tag = true
						} else if board[nx][ny] == 'O' && !vis[nx][ny] {
							vis[nx][ny] = true
							list = append(list, []int{nx, ny})
							tmp = append(tmp, []int{nx, ny})
						}
					}
				}
				now = tmp
			}

			if !tag {
				for l := range list {
					board[list[l][0]][list[l][1]] = 'X'
				}
			}

		}
	}
}
