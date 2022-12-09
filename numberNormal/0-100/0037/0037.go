package leetcode

func solveSudoku(board [][]byte) {
	var row [9][9]bool
	var column [9][9]bool
	var squares [3][3][9]bool

	for i := range board {
		for j := range board[0] {

			if board[i][j] == '.' {
				continue
			}

			t := int(board[i][j] - '1')

			row[i][t], column[j][t], squares[i/3][j/3][t] = true, true, true
		}
	}

	var dfs func(int, int) bool

	dfs = func(i int, j int) bool {
		if j == 9 {
			i, j = i+1, 0
		}

		if i == 9 {
			return true
		}

		if board[i][j] != '.' {
			return dfs(i, j+1)
		} else {
			for t := 0; t < 9; t++ {
				if row[i][t] || column[j][t] || squares[i/3][j/3][t] {
					continue
				}

				board[i][j] = byte(t + '1')
				row[i][t], column[j][t], squares[i/3][j/3][t] = true, true, true

				res := dfs(i, j+1)

				if !res {
					board[i][j] = '.'
					row[i][t], column[j][t], squares[i/3][j/3][t] = false, false, false
				} else {
					return true
				}

			}
			return false
		}
	}

	dfs(0, 0)
}
