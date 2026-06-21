package level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}

	res := make([][]int, 0)

	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) != 0 {
		qLen := len(q)
		innerRes := make([]int, 0)

		for i := 0; i < qLen; i++ {
			node := q[0]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			innerRes = append(innerRes, node.Val)
			q = q[1:]
		}
		res = append(res, innerRes)

	}

	return res
}
