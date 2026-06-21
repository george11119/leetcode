package subtree_of_another_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) != 0 {
		node := q[0]

		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if isSameTree(node, subRoot) {
			return true
		}

		q = q[1:]
	}

	return false
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q == nil {
		return false
	} else if p == nil && q != nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	leftSubtreeSame := isSameTree(p.Left, q.Left)
	rightSubtreeSame := isSameTree(p.Right, q.Right)

	return leftSubtreeSame && rightSubtreeSame
}
