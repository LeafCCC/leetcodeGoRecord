// go中默认的int根据系统位数决定 leetcode里都是64位 所以回避了很多溢出问题 不太合理
package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	y := 0
	t := x

	for t > 0 {
		y *= 10
		y += t % 10
		t /= 10
	}

	return y == x
}
