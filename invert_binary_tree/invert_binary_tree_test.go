package invert_binary_tree

import (
	"math"
	"reflect"
	"testing"
)

const NULL = math.MinInt

func createBinaryTree(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		if v != NULL {
			nodes[i] = &TreeNode{Val: v}
		}
	}

	for i, node := range nodes {
		if node == nil {
			continue
		}
		if left := 2*i + 1; left < len(nodes) {
			node.Left = nodes[left]
		}
		if right := 2*i + 2; right < len(nodes) {
			node.Right = nodes[right]
		}
	}

	return nodes[0]
}

func binaryTreeToSlice(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	vals := make(map[int]int)
	maxIndex := 0
	var traverse func(node *TreeNode, i int)
	traverse = func(node *TreeNode, i int) {
		if node == nil {
			return
		}
		vals[i] = node.Val
		if i > maxIndex {
			maxIndex = i
		}
		traverse(node.Left, 2*i+1)
		traverse(node.Right, 2*i+2)
	}
	traverse(root, 0)

	size := 1
	for size <= maxIndex {
		size = size*2 + 1
	}

	result := make([]int, size)
	for i := range result {
		if v, ok := vals[i]; ok {
			result[i] = v
		} else {
			result[i] = NULL
		}
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
		bTreeVals := []int{1, NULL, 3, NULL, NULL, 2, NULL}
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

func TestInvertTree(t *testing.T) {
	tt := []struct {
		bTreeVals []int
		want      []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 3, 2, 7, 6, 5, 4},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{2, 1, 3},
			[]int{2, 3, 1},
		},
		{
			[]int{2, 3, NULL},
			[]int{2, NULL, 3},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			bTree := createBinaryTree(tc.bTreeVals)
			bTree = invertTree(bTree)

			got := binaryTreeToSlice(bTree)
			want := tc.want

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v want %v", got, want)
			}
		})
	}
}
