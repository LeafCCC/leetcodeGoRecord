package leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	q := []*Node{root}
	for len(q) != 0 {
		tmp := []*Node{}
		for i := 0; i < len(q)-1; i++ {
			q[i].Next = q[i+1]
			if q[i].Left != nil {
				tmp = append(tmp, q[i].Left, q[i].Right)
			}
		}

		if q[len(q)-1].Left != nil {
			tmp = append(tmp, q[len(q)-1].Left, q[len(q)-1].Right)
		}
		q = tmp
	}
	return root
}
