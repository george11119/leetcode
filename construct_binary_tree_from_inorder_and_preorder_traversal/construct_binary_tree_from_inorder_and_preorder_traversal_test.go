package construct_binary_tree_from_inorder_and_preorder_traversal

import (
	"math"
	"reflect"
	"testing"
)

const NULL = math.MinInt

func createBinaryTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == NULL {
		return nil
	}

	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		if i < len(vals) {
			if vals[i] != NULL {
				node.Left = &TreeNode{Val: vals[i]}
				queue = append(queue, node.Left)
			}
			i++
		}
		if i < len(vals) {
			if vals[i] != NULL {
				node.Right = &TreeNode{Val: vals[i]}
				queue = append(queue, node.Right)
			}
			i++
		}
	}

	return root
}

func binaryTreeToSlice(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, NULL)
			continue
		}

		result = append(result, node.Val)
		queue = append(queue, node.Left, node.Right)
	}

	// Trim trailing NULLs left behind by leaf children.
	for len(result) > 0 && result[len(result)-1] == NULL {
		result = result[:len(result)-1]
	}

	return result
}

func TestBinaryTreeUtils(t *testing.T) {
	t.Run("empty tree edge case", func(t *testing.T) {
		bTreeVals := make([]int, 0)
		root := createBinaryTree(bTreeVals)

		if root != nil {
			t.Fatalf("Binary tree was formed wrong")
		}

		got := binaryTreeToSlice(root)

		if !reflect.DeepEqual(got, bTreeVals) {
			t.Fatalf("got %v want %v", got, bTreeVals)
		}
	})

	t.Run("binary tree with null", func(t *testing.T) {
		bTreeVals := []int{1, NULL, 3, 2}
		root := createBinaryTree(bTreeVals)

		if root.Val != 1 ||
			root.Left != nil || root.Right.Val != 3 ||
			root.Right.Left.Val != 2 || root.Right.Right != nil {
			t.Fatalf("Binary tree was formed wrong")
		}

		got := binaryTreeToSlice(root)

		if !reflect.DeepEqual(got, bTreeVals) {
			t.Fatalf("got %v want %v", got, bTreeVals)
		}
	})

	t.Run("binary tree", func(t *testing.T) {
		bTreeVals := []int{1, 2, 3, 4, 5, 6, 7}
		root := createBinaryTree(bTreeVals)

		if root.Val != 1 ||
			root.Left.Val != 2 || root.Right.Val != 3 ||
			root.Left.Left.Val != 4 || root.Left.Right.Val != 5 ||
			root.Right.Left.Val != 6 || root.Right.Right.Val != 7 ||
			root.Left.Right.Left != nil || root.Left.Right.Right != nil ||
			root.Right.Right.Left != nil || root.Right.Right.Right != nil {
			t.Fatalf("Binary tree was formed wrong")
		}

		got := binaryTreeToSlice(root)

		if !reflect.DeepEqual(got, bTreeVals) {
			t.Fatalf("got %v want %v", got, bTreeVals)
		}
	})
}

func Test(t *testing.T) {
	tt := []struct {
		preorder []int
		inorder  []int
		want     []int
	}{
		{
			[]int{4, 2, 1, 3, 6, 5, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{4, 2, 6, 1, 3, 5, 7},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{2, 1, 3, 4},
			[]int{1, 2, 3, NULL, NULL, NULL, 4},
		},
		{
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
			[]int{3, 9, 20, NULL, NULL, 15, 7},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			gotTree := buildTree(tc.preorder, tc.inorder)

			got := binaryTreeToSlice(gotTree)
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
