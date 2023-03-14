package leetcode

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	var dfs func(int, int, int) bool

	dfs = func(level int, x, y int) bool {
		if level == len(word) {
			return true
		}

		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != word[level] || vis[x][y] {
			return false
		}

		vis[x][y] = true

		defer func() {
			vis[x][y] = false
		}()

		return dfs(level+1, x+1, y) || dfs(level+1, x, y+1) || dfs(level+1, x, y-1) || dfs(level+1, x-1, y)
	}

	for i := range board {
		for j := range board[i] {
			if dfs(0, i, j) {
				return true
			}
		}
	}

	return false
}
