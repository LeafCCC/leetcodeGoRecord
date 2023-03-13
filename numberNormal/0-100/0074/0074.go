package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if matrix[0][0] > target {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	begin, end := 0, m
	for begin < end-1 {
		mid := begin + (end-begin)/2
		if matrix[mid][0] <= target {
			begin = mid
		} else {
			end = mid
		}
	}

	row := begin
	begin, end = -1, n
	for begin < end-1 {
		mid := begin + (end-begin)/2
		if matrix[row][mid] < target {
			begin = mid
		} else if matrix[row][mid] > target {
			end = mid
		} else {
			return true
		}
	}

	return false
}
