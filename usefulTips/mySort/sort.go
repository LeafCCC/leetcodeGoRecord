package main

import (
	"fmt"
	"math/rand"
)

// 冒泡排序——O(n²)
func BubbleSort(nums []int) {
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		swapped := false
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// 插入排序——O(n²)
func InsertionSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			} else {
				break
			}
		}
	}
}

// 合并排序——O(nlgn)算法
func MergeSort(nums []int, begin, end int) {
	if begin >= end {
		return
	}

	mid := begin + (end-begin)/2

	MergeSort(nums, begin, mid)
	MergeSort(nums, mid+1, end)
	Merge(nums, mid, begin, end)
}

func Merge(nums []int, mid, begin, end int) {
	if begin > mid || mid >= end {
		return
	}

	tmp := make([]int, end-begin+1)
	t1, t2 := begin, mid+1
	now := 0

	for t1 < mid+1 || t2 < end+1 {
		if t1 == mid+1 {
			for i := t2; i < end+1; i++ {
				tmp[now] = nums[i]
				now++
			}
			t2 = end + 1
		} else if t2 == end+1 {
			for i := t1; i < mid+1; i++ {
				tmp[now] = nums[i]
				now++
			}
			t1 = mid + 1
		} else if nums[t1] < nums[t2] {
			tmp[now] = nums[t1]
			t1++
		} else {
			tmp[now] = nums[t2]
			t2++
		}

		now++
	}

	for i := begin; i <= end; i++ {
		nums[i] = tmp[i-begin]
	}
}

// 快速排序——平均O(nlgn)算法 最坏n²
// func QuickSort(nums []int, begin, end int) {
// 	if begin < end {
// 		pivot := Partition(nums, begin, end)
// 		QuickSort(nums, begin, pivot-1)
// 		QuickSort(nums, pivot+1, end)
// 	}
// }

// func Partition(nums []int, begin, end int) int {
// 	r := rand.Intn(end-begin+1) + begin
// 	nums[r], nums[end] = nums[end], nums[r]
// 	target := nums[end]
// 	where := begin

// 	for i := begin; i < end; i++ {
// 		if nums[i] < target {
// 			nums[i], nums[where] = nums[where], nums[i]
// 			where++
// 		}
// 	}

//		nums[end], nums[where] = nums[where], nums[end]
//		return where
//	}
func QuickSort(nums []int, begin, end int) {
	if begin < end {
		pivot := Partition(nums, begin, end)
		QuickSort(nums, begin, pivot-1)
		QuickSort(nums, pivot+1, end)
	}
}

func Partition(nums []int, begin, end int) int {
	r := rand.Intn(end-begin+1) + begin
	nums[r], nums[end] = nums[end], nums[r]
	target := nums[end]
	where := begin

	for i := begin; i < end; i++ {
		if target > nums[i] {
			nums[i], nums[where] = nums[where], nums[i]
			where++
		}
	}

	nums[end], nums[where] = nums[where], nums[end]
	return where
}

func main() {
	fmt.Println("冒泡排序")
	nums := []int{2, 4, 5, 2, 3, 5, 100, 4, 1, 5, 7, 24, 24234, 24, 24, 4, 43, 65, 86, 65, 3, 2, 2}
	fmt.Println(nums)
	BubbleSort(nums)
	fmt.Println(nums)
	fmt.Println()

	fmt.Println("插入排序")
	nums = []int{2, 4, 5, 2, 3, 5, 100, 4, 1, 5, 7, 24, 24234, 24, 24, 4, 43, 65, 86, 65, 3, 2, 2}
	fmt.Println(nums)
	InsertionSort(nums)
	fmt.Println(nums)
	fmt.Println()

	fmt.Println("归并排序")
	nums = []int{2, 4, 5, 2, 3, 5, 100, 4, 1, 5, 7, 24, 24234, 24, 24, 4, 43, 65, 86, 65, 3, 2, 2}
	fmt.Println(nums)
	MergeSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	fmt.Println()

	fmt.Println("快速排序")
	nums = []int{2, 4, 5, 2, 3, 5, 100, 4, 1, 5, 7, 24, 24234, 24, 24, 4, 43, 65, 86, 65, 3, 2, 2}
	fmt.Println(nums)
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	fmt.Println()

}
