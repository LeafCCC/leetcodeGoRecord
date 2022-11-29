package leetcode

import "sort"

func threeSum(nums []int) (res [][]int) {
	sort.Ints(nums)

	for i := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		target := -nums[i]
		if target < 0 {
			break
		}

		begin := i + 1
		end := len(nums) - 1
		for begin < end {
			tmp := nums[begin] + nums[end]
			if tmp == target {
				res = append(res, []int{nums[i], nums[begin], nums[end]})
				begin++
				end--
			} else if tmp > target {
				end--
			} else {
				begin++
			}

			for begin > i+1 && begin < len(nums)-1 && nums[begin] == nums[begin-1] {
				begin++
			}

			for end < len(nums)-1 && end > 0 && nums[end] == nums[end+1] {
				end--
			}
		}
	}
	return

}
