package construct_binary_tree_from_inorder_and_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	indices := make(map[int]int)
	for i, v := range inorder {
		indices[v] = i
	}

	preIdx := 0

	var dfs func(left, right int) *TreeNode
	dfs = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		node := &TreeNode{Val: preorder[preIdx]}
		mid := indices[preorder[preIdx]]
		preIdx++

		node.Left = dfs(left, mid-1)
		node.Right = dfs(mid+1, right)

		return node
	}

	return dfs(0, len(preorder)-1)
}
