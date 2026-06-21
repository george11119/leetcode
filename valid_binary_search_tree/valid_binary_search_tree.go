package valid_binary_search_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt, math.MaxInt)
}

func helper(root *TreeNode, lowerBound int, upperBound int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lowerBound || root.Val >= upperBound {
		return false
	}

	return helper(root.Left, lowerBound, root.Val) && helper(root.Right, root.Val, upperBound)
}
