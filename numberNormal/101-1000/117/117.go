/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
package leetcode

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	q := []*Node{root}
	for len(q) != 0 {
		tmp := []*Node{}

		for i := 0; i < len(q); i++ {
			if i < len(q)-1 {
				q[i].Next = q[i+1]
			}

			if q[i].Left != nil {
				tmp = append(tmp, q[i].Left)
			}
			if q[i].Right != nil {
				tmp = append(tmp, q[i].Right)
			}
		}

		q = tmp
	}

	return root
}
