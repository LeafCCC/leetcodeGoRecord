package leetcode

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) (ret [][]int) {
	items := append(items1, items2...)

	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	ret = append(ret, []int{items[0][0], items[0][1]})

	for _, item := range items[1:] {
		if ret[len(ret)-1][0] == item[0] {
			ret[len(ret)-1][1] += item[1]
		} else {
			ret = append(ret, []int{item[0], item[1]})
		}
	}

	return ret
}
