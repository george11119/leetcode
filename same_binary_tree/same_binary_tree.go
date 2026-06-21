package same_binary_tree

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
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
