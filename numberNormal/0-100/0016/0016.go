package leetcode

import "sort"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := abs(nums[0] + nums[1] + nums[2] - target)
	theSum := nums[0] + nums[1] + nums[2]

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		begin := i + 1
		end := len(nums) - 1

		for begin < end {
			t := nums[i] + nums[begin] + nums[end]
			if t == target {
				return target
			} else if t > target {
				end--

				for end > 0 && nums[end] == nums[end+1] {
					end--
				}
			} else {
				begin++

				for begin < len(nums) && nums[begin] == nums[begin-1] {
					begin++
				}
			}

			if abs(t-target) < res {
				res = abs(t - target)
				theSum = t
			}
		}
	}
	return theSum
}
