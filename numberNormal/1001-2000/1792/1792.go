package leetcode

import "container/heap"

type myHeap [][]int

func (h myHeap) Len() int { return len(h) }
func (h myHeap) Less(i, j int) bool {
	return (h[i][1]-h[i][0])*(h[j][1]*h[j][1]+h[j][1]) > (h[j][1]-h[j][0])*(h[i][1]*h[i][1]+h[i][1])
}
func (h myHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *myHeap) Push(x interface{})   {}
func (h *myHeap) Pop() (_ interface{}) { return }

func maxAverageRatio(classes [][]int, extraStudents int) (ans float64) {
	h := myHeap(classes)
	heap.Init(&h)
	for i := 0; i < extraStudents; i++ {
		h[0][1]++
		h[0][0]++
		heap.Fix(&h, 0)
	}

	for _, x := range h {
		ans += float64(x[0]) / float64(x[1])
	}

	return ans / float64(len(h))
}
