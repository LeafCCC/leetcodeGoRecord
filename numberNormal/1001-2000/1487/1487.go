package leetcode

import "strconv"

func getFolderNames(names []string) []string {
	record := make(map[string]int)

	res := make([]string, len(names))

	for i := range names {
		if record[names[i]] == 0 {
			res[i] = names[i]
		} else {
			tmp := names[i] + "(" + strconv.Itoa(record[names[i]]) + ")"
			for record[tmp] != 0 {
				record[names[i]]++
				tmp = names[i] + "(" + strconv.Itoa(record[names[i]]) + ")"
			}
			res[i] = tmp
			record[tmp]++
		}
		record[names[i]]++
	}

	return res
}
