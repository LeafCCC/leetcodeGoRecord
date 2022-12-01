package leetcode

//left记录左括号的数量 right记录右括号数量
func f(left int, right int, n int, tmp []byte, res *[]string) {
	if right == n {
		*res = append(*res, string(tmp))
	} else {
		if left < n {
			f(left+1, right, n, append(tmp, '('), res)
		}

		if left > right {
			f(left, right+1, n, append(tmp, ')'), res)
		}
	}

}

func generateParenthesis(n int) (res []string) {
	f(0, 0, n, []byte{}, &res)
	return
}
