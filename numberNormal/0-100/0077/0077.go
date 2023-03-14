package leetcode

func combine(n int, k int) [][]int {
	res := [][]int{}
	tmp := []int{}

	var backTrack func(int, int)
	backTrack = func(now int, level int) {
		if level == k {
			res = append(res, append([]int{}, tmp...))
			return
		}
		for i := now; i <= n-k+level+1; i++ {
			tmp = append(tmp, i)
			backTrack(i+1, level+1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	backTrack(1, 0)
	return res
}
