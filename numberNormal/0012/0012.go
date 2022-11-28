package leetcode

import "sort"

//用结构体更好 省去了排序
func intToRoman(num int) string {
	rcd := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	keys := make([]int, len(rcd))
	i := 0
	for key := range rcd {
		keys[i] = key
		i++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	res := []byte{}
	for _, key := range keys {
		for num >= key {
			num -= key
			res = append(res, rcd[key]...)
		}
		if num == 0 {
			break
		}
	}

	return string(res)
}
