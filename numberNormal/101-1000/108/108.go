package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)

	if n == 0 {
		return nil
	}

	target := int(n / 2)

	root := &TreeNode{nums[target], nil, nil}

	root.Left = sortedArrayToBST(nums[0:target])
	root.Right = sortedArrayToBST(nums[target+1:])

	return root
}
