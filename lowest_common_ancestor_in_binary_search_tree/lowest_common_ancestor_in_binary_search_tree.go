package lowest_common_ancestor_in_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//// recursive
//func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
//	if root == nil || p == nil || q == nil {
//		return nil
//	}
//
//	if p.Val < root.Val && q.Val < root.Val {
//		return lowestCommonAncestor(root.Left, p, q)
//	} else if p.Val > root.Val && q.Val > root.Val {
//		return lowestCommonAncestor(root.Right, p, q)
//	} else {
//		return root
//	}
//}

// iterative
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	cur := root

	for cur != nil {
		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}

	return nil
}
