package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	now1 := m - 1
	now2 := n - 1

	for i := m + n - 1; i >= 0; i-- {
		if now2 < 0 {
			return
		} else if now1 < 0 || nums1[now1] <= nums2[now2] {
			nums1[i] = nums2[now2]
			now2--
		} else if nums1[now1] > nums2[now2] {
			nums1[i] = nums1[now1]
			now1--
		}
	}

	return
}
