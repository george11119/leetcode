package kth_smallest_integer_in_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	res := -1
	_, res = helper(root, k, res)
	return res
}

func helper(root *TreeNode, k int, res int) (int, int) {
	if root == nil {
		return k, res
	}

	k, res = helper(root.Left, k, res)
	k--
	if k == 0 {
		res = root.Val
	}
	k, res = helper(root.Right, k, res)

	return k, res
}
