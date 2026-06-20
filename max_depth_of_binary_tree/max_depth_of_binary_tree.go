package max_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive
//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//
//	return max(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
//}

// iterative with BFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := make([]*TreeNode, 0)
	depth := 0
	q = append(q, root)

	for len(q) != 0 {
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			node := q[0]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			q = q[1:]
		}

		depth++
	}

	return depth
}
