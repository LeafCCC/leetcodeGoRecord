package leetcode

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	record := make(map[int]*Node)
	vis := make(map[*Node]bool)
	now := []*Node{node}

	res := &Node{Val: node.Val}
	record[node.Val] = res
	vis[node] = true

	for len(now) != 0 {
		tmp := []*Node{}

		for _, n := range now {

			target := record[n.Val]

			for _, neighbor := range n.Neighbors {
				if !vis[neighbor] {
					tmp = append(tmp, neighbor)
					vis[neighbor] = true
				}

				if record[neighbor.Val] != nil {
					tNeighbor := record[neighbor.Val]
					target.Neighbors = append(target.Neighbors, tNeighbor)
				} else {
					tNeighbor := &Node{Val: neighbor.Val}
					record[neighbor.Val] = tNeighbor
					target.Neighbors = append(target.Neighbors, tNeighbor)
				}
			}
		}

		now = tmp
	}

	return res
}
