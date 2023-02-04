package leetcode

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	type pair struct{ x, color int }
	record := make([][]pair, n)

	//红0 蓝1
	//record[0]为例 代表从0出发 到x 颜色是color
	for _, e := range redEdges {
		record[e[0]] = append(record[e[0]], pair{e[1], 0})
	}

	for _, e := range blueEdges {
		record[e[0]] = append(record[e[0]], pair{e[1], 1})
	}

	//初始化answer
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	//初始化队列
	q := []pair{pair{0, 0}, pair{0, 1}}

	//初始化访问队列
	vis := make([][2]bool, n)

	//初始化距离
	d := 0

	for len(q) != 0 {
		for k := len(q); k > 0; k-- {
			node := q[0]
			q = q[1:]

			x, c := node.x, node.color
			if ans[x] == -1 {
				ans[x] = d
			}

			vis[x][c] = true

			c ^= 1

			for _, r := range record[x] {
				if r.color == c && !vis[r.x][c] {
					q = append(q, pair{r.x, c})
				}
			}
		}
		d++
	}

	return ans
}
