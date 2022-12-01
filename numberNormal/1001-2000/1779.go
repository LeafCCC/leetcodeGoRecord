package leetcode

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func caculateD(a []int, b []int) int {
	if a[0] == b[0] {
		return abs(a[1] - b[1])
	}
	return abs(a[0] - b[0])
}

func nearestValidPoint(x int, y int, points [][]int) int {
	res := -1
	distance := 0
	me := []int{x, y}
	for i := range points {
		if points[i][0] == x || points[i][1] == y {
			if distance > caculateD(points[i], me) || res == -1 {
				res = i
				distance = caculateD(points[i], me)
			}
		}
	}
	return res
}
