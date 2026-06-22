package delete_node_in_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMin(node *TreeNode) int {
	cur := node
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur.Val
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	// not found
	if root == nil {
		return root
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil && root.Right != nil {
			return root.Right
		} else if root.Left != nil && root.Right == nil {
			return root.Left
		} else {
			minInRight := findMin(root.Right)
			root.Val = minInRight
			root.Right = deleteNode(root.Right, minInRight)
		}
	}

	return root
}
