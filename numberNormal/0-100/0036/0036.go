package leetcode

func isValidSudoku(board [][]byte) bool {
	var row [9][9]bool
	var column [9][9]bool
	var squares [3][3][9]bool

	for i := range board {
		for j := range board[0] {

			if board[i][j] == '.' {
				continue
			}

			t := int(board[i][j] - '1')

			if row[i][t] || column[j][t] || squares[i/3][j/3][t] {
				return false
			}

			row[i][t], column[j][t], squares[i/3][j/3][t] = true, true, true
		}
	}

	return true
}
