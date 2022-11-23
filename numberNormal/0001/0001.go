// 1 两数之和
// 直接用一个哈希表解决即可
package leetcode

func twoSum(nums []int, target int) []int {
	record := make(map[int]int)
	for i, val := range nums {
		if idx, ok := record[target-val]; ok {
			return []int{i, idx}
		} else {
			record[val] = i
		}
	}
	return nil

}
