package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// original nodes as key, cloned node's pointers as the value
	graphMap := make(map[*Node]*Node)
	graphMap[node] = &Node{node.Val, make([]*Node, 0)}

	// queue contains original nodes
	q := make([]*Node, 0)
	q = append(q, node)

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		for _, n := range cur.Neighbors {
			if _, exists := graphMap[n]; !exists {
				graphMap[n] = &Node{n.Val, make([]*Node, 0)}
				q = append(q, n)
			}

			graphMap[cur].Neighbors = append(graphMap[cur].Neighbors, graphMap[n])
		}
	}

	return graphMap[node]
}
