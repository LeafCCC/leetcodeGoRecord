package leetcode

func validPath(n int, edges [][]int, source int, destination int) bool {
	ds := make([]int, n)

	for i := range ds {
		ds[i] = i
	}

	var find func(a int) int
	find = func(a int) int {
		if ds[a] != a {
			ds[a] = find(ds[a])
		}

		return ds[a]
	}

	merge := func(a, b int) {
		ds[find(a)] = find(b)
	}

	for _, edge := range edges {
		merge(edge[0], edge[1])
	}

	return find(source) == find(destination)
}
