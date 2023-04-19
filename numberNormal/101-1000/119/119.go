package leetcode

func getRow(rowIndex int) []int {
	now := []int{1}

	for i := 1; i <= rowIndex; i++ {
		tmp := []int{1}
		for j := 0; j < i-1; j++ {
			tmp = append(tmp, now[j]+now[j+1])
		}
		tmp = append(tmp, 1)
		now = tmp
	}

	return now
}
