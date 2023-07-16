
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	//记录旧到新的对应
	old2new := make(map[*Node]*Node)

	res := &Node{Val: head.Val}
	now2 := res
	now1 := head.Next

	old2new[head] = res

	for now1 != nil {
		tmp := &Node{Val: now1.Val}
		now2.Next = tmp
		now2 = now2.Next

		old2new[now1] = now2

		now1 = now1.Next
	}

	old2new[now1] = now2.Next

	now1 = head
	now2 = res
	for now1 != nil {
		now2.Random = old2new[now1.Random]
		now1 = now1.Next
		now2 = now2.Next
	}

	return res
}