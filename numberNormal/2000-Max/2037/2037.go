package leetcode

import "sort"

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	abs := func(a int) int {
		if a > 0 {
			return a
		}
		return -a
	}

	res := 0

	for i := range seats {
		res += abs(seats[i] - students[i])
	}

	return res
}
