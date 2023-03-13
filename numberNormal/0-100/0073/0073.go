package leetcode

func setZeroes(matrix [][]int) {
	row, column := map[int]bool{}, map[int]bool{}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				row[i] = true
				column[j] = true
			}
		}
	}

	for i, _ := range row {
		for j := range matrix[i] {
			matrix[i][j] = 0
		}
	}

	for j, _ := range column {
		for i := range matrix {
			matrix[i][j] = 0
		}
	}
}
