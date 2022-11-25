package leetcode

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// 找到两个有序数组中第k小的数
func findK(nums1 []int, nums2 []int, k int) int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 {
		return nums2[k-1]
	}
	if n2 == 0 {
		return nums1[k-1]
	}

	if k == 1 {
		return min(nums1[0], nums2[0])
	}

	t := k / 2
	if n1 < t {
		if nums1[n1-1] > nums2[t-1] {
			return findK(nums1, nums2[t:], k-t)
		} else {
			return nums2[k-n1-1]
		}
	} else if n2 < t {
		if nums1[t-1] <= nums2[n2-1] {
			return findK(nums1[t:], nums2, k-t)
		} else {
			return nums1[k-n2-1]
		}
	}
	if nums1[t-1] < nums2[t-1] {
		return findK(nums1[t:], nums2, k-t)
	}
	return findK(nums1, nums2[t:], k-t)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)

	a := findK(nums1, nums2, (n+1)/2)
	b := findK(nums1, nums2, n/2+1)
	return (float64(a) + float64(b)) / 2
}
