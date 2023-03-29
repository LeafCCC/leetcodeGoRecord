package leetcode

func help(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := int(len(nums) / 2)
	root := &TreeNode{Val: nums[mid]}

	root.Left = help(nums[:mid])
	root.Right = help(nums[mid+1:])
	return root

}
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	nums := []int{}

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	return help(nums)

}
