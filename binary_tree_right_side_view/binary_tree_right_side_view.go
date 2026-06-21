package binary_tree_right_side_view

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) != 0 {
		qLen := len(q)

		for i := 0; i < qLen; i++ {
			node := q[0]
			if i == qLen-1 {
				res = append(res, node.Val)
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}

			q = q[1:]
		}
	}

	return res
}
