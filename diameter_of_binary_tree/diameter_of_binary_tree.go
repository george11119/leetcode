package diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, res := dfs(root)
	return res
}

func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftH, leftD := dfs(root.Left)
	rightH, rightD := dfs(root.Right)

	height := 1 + max(leftH, rightH)
	diameter := max(leftH+rightH, max(leftD, rightD))

	return height, diameter
}
