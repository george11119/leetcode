package reorder_linked_list

import (
	"testing"
)

func createLinkedListFromSlice(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := ListNode{arr[0], nil}

	tmp := &head
	for i := 1; i < len(arr); i++ {
		tmp.Next = &ListNode{arr[i], nil}
		tmp = tmp.Next
	}

	return &head
}

func compareLinkedListValues(head *ListNode, arr []int) bool {
	tmp := head
	for i := 0; i < len(arr); i++ {
		if tmp == nil {
			return false
		}
		if arr[i] != tmp.Val {
			return false
		}
		tmp = tmp.Next
	}

	return tmp == nil
}

func linkedListToSlice(head *ListNode) []int {
	res := make([]int, 0)

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

func TestReorderList(t *testing.T) {
	tt := []struct {
		l    []int
		want []int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5, 6},
			[]int{0, 6, 1, 5, 2, 4, 3},
		},
		{
			[]int{2, 4, 6, 8},
			[]int{2, 8, 4, 6},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			ll := createLinkedListFromSlice(tc.l)
			reorderList(ll)

			if !compareLinkedListValues(ll, tc.want) {
				t.Fatalf("got %v, want %v", linkedListToSlice(ll), tc.want)
			}
		})
	}
}
