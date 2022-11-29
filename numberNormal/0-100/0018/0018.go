package leetcode

import "sort"

func fourSum(nums []int, target int) (res [][]int) {
	sort.Ints(nums)

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 && nums[i] >= target {
			break
		}

		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			sum := nums[i] + nums[j]

			if nums[j] > 0 && sum >= target {
				break
			}

			begin := j + 1
			end := len(nums) - 1

			for begin < end {
				sum = nums[i] + nums[j] + nums[begin] + nums[end]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[begin], nums[end]})
					begin++
					end--
				} else if sum > target {
					end--
				} else {
					begin++
				}

				for begin > j+1 && begin < len(nums)-1 && nums[begin] == nums[begin-1] {
					begin++
				}
				for end < len(nums)-1 && end > j+1 && nums[end] == nums[end+1] {
					end--
				}
			}
		}
	}
	return res
}
