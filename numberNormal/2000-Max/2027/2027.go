package leetcode

func minimumMoves(s string) int {
	res := 0
	i := 0
	for i < len(s) {
		if s[i] == 'X' {
			i += 3
			res++
		} else {
			i++
		}
	}
	return res
}
