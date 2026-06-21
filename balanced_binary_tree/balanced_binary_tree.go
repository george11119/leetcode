package balanced_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func helper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftHeight, leftBalanced := helper(root.Left)
	rightHeight, rightBalanced := helper(root.Right)

	balanced := leftBalanced && rightBalanced && abs(leftHeight-rightHeight) <= 1

	return 1 + max(leftHeight, rightHeight), balanced
}

func isBalanced(root *TreeNode) bool {
	_, balanced := helper(root)
	return balanced
}
