package count_good_nodes_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return helper(root, root.Val)
}

func helper(root *TreeNode, val int) int {
	if root == nil {
		return 0
	}

	if root.Val >= val {
		return 1 + helper(root.Left, root.Val) + helper(root.Right, root.Val)
	}
	return helper(root.Left, val) + helper(root.Right, val)
}
