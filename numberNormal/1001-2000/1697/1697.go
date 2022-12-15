package leetcode

import "sort"

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	//并查集初始化
	ds := make([]int, n)
	for i := range ds {
		ds[i] = i
	}

	//查询
	var find func(int) int
	find = func(x int) int {
		if ds[x] != x {
			ds[x] = find(ds[x])
		}
		return ds[x]
	}

	//合并
	merge := func(from, to int) {
		ds[find(from)] = find(to)
	}

	//离线查询
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	//记录queries下标
	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][2] < queries[j][2]
	})

	ans := make([]bool, len(queries))
	k := 0 //记录边的下标

	for i := 0; i < len(queries); i++ {
		for ; k < len(edgeList) && edgeList[k][2] < queries[i][2]; k++ {
			merge(edgeList[k][0], edgeList[k][1])
		}
		ans[queries[i][3]] = find(queries[i][0]) == find(queries[i][1])
	}

	return ans
}
