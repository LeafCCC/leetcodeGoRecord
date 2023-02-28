package leetcode

import (
	"container/heap"
	"math"
)

func minimumTime(grid [][]int) int {
	dirs := []int{-1, 0, 1, 0, -1}

	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}

	m, n := len(grid), len(grid[0])
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	dis[0][0] = 0

	h := &hp{point{0, 0, 0}}

	for {
		p := heap.Pop(h).(point)
		d, x, y := p.d, p.x, p.y

		if x == m-1 && y == n-1 {
			return d
		}

		for i := 0; i < 4; i++ {
			nx, ny := x+dirs[i], y+dirs[i+1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				nd := max(d+1, grid[nx][ny])
				nd += (nd - nx - ny) % 2
				if nd < dis[nx][ny] {
					dis[nx][ny] = nd
					heap.Push(h, point{nd, nx, ny})
				}
			}
		}
	}

}

type point struct{ d, x, y int }
type hp []point

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(point)) }
func (h *hp) Pop() (y interface{}) {
	y = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
