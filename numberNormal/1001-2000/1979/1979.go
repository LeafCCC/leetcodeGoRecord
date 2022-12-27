package leetcode

import "sort"

func findGCD(nums []int) int {
	sort.Ints(nums)
	return gcd(nums[0], nums[len(nums)-1])
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
