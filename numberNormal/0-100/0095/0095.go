package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func generateTrees(n int) []*TreeNode {
    
    var helper func(int, int) []*TreeNode
    
    helper = func(start, end int) (res []*TreeNode){
        if start > end{
            return []*TreeNode{nil}
        }

        for i := start; i <= end; i++{
            l := helper(start, i-1)
            r := helper(i+1, end)
            for _, ln := range l{
                for _, rn := range r{
                    tmp := &TreeNode{Val: i}
                    tmp.Left = ln
                    tmp.Right = rn
                    res = append(res, tmp)
                }
            }
        }
        return
    }
    
    return helper(1, n)
}