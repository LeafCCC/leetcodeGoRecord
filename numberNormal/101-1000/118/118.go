package leetcode

func generate(numRows int) [][]int {
	res := [][]int{[]int{1}}

	for i := 2; i <= numRows; i++ {
		tmp := []int{res[i-2][0]}
		for j := 0; j < len(res)-1; j++ {
			tmp = append(tmp, res[i-2][j]+res[i-2][j+1])
		}
		tmp = append(tmp, res[i-2][len(res)-1])
		res = append(res, tmp)
	}

	return res
}
