package clone_graph

import (
	"reflect"
	"testing"
)

// buildGraph constructs a graph from a 1-indexed adjacency list and returns
// the first node (value 1), or nil if the list is empty.
func buildGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}

	nodes := make([]*Node, len(adjList))
	for i := range adjList {
		nodes[i] = &Node{Val: i + 1}
	}

	for i, neighbors := range adjList {
		for _, v := range neighbors {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[v-1])
		}
	}

	return nodes[0]
}

// toAdjList walks the graph starting at node and returns its 1-indexed
// adjacency list representation.
func toAdjList(node *Node) [][]int {
	if node == nil {
		return [][]int{}
	}

	seen := make(map[int][]int)
	q := []*Node{node}
	seen[node.Val] = nil

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		neighbors := make([]int, 0, len(cur.Neighbors))
		for _, n := range cur.Neighbors {
			neighbors = append(neighbors, n.Val)
			if _, found := seen[n.Val]; !found {
				seen[n.Val] = nil
				q = append(q, n)
			}
		}
		seen[cur.Val] = neighbors
	}

	adjList := make([][]int, len(seen))
	for val, neighbors := range seen {
		adjList[val-1] = neighbors
	}
	return adjList
}

// collectNodes returns every node reachable from the given node.
func collectNodes(node *Node) []*Node {
	if node == nil {
		return nil
	}

	seen := make(map[*Node]bool)
	q := []*Node{node}
	seen[node] = true

	var nodes []*Node
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		nodes = append(nodes, cur)

		for _, n := range cur.Neighbors {
			if !seen[n] {
				seen[n] = true
				q = append(q, n)
			}
		}
	}
	return nodes
}

func TestCloneGraph(t *testing.T) {
	tt := []struct {
		name    string
		adjList [][]int
	}{
		{
			name:    "three connected nodes",
			adjList: [][]int{{2}, {1, 3}, {2}},
		},
		// {
		// 	name:    "single node no neighbors",
		// 	adjList: [][]int{{}},
		// },
		// {
		// 	name:    "empty graph",
		// 	adjList: [][]int{},
		// },
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			original := buildGraph(tc.adjList)
			clone := cloneGraph(original)

			// The clone must reproduce the same adjacency list.
			got := toAdjList(clone)
			if !reflect.DeepEqual(got, tc.adjList) {
				t.Fatalf("got %v want %v", got, tc.adjList)
			}

			// The clone must be a deep copy: no node pointer may be shared
			// between the original graph and its clone.
			originalNodes := make(map[*Node]bool)
			for _, n := range collectNodes(original) {
				originalNodes[n] = true
			}
			for _, n := range collectNodes(clone) {
				if originalNodes[n] {
					t.Fatalf("clone shares node pointer with original for val %d", n.Val)
				}
			}
		})
	}
}
