package leetcode

import "sort"

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	rcd := make(map[int]int)

	sort.Ints(nums1)
	sort.Ints(nums2)
	sort.Ints(nums3)

	rcd[nums1[0]]++
	rcd[nums2[0]]++
	rcd[nums3[0]]++

	for i := 1; i < len(nums1); i++ {
		if nums1[i] != nums1[i-1] {
			rcd[nums1[i]]++
		}
	}

	for i := 1; i < len(nums2); i++ {
		if nums2[i] != nums2[i-1] {
			rcd[nums2[i]]++
		}
	}

	for i := 1; i < len(nums3); i++ {
		if nums3[i] != nums3[i-1] {
			rcd[nums3[i]]++
		}
	}

	res := []int{}
	for i, val := range rcd {
		if val > 1 {
			res = append(res, i)
		}
	}

	return res
}
