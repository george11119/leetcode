package reverse_linked_list

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

func TestReverseList(t *testing.T) {
	tt := []struct {
		input []int
		want  []int
	}{
		{
			[]int{0, 1, 2, 3},
			[]int{3, 2, 1, 0},
		},
		//{
		//	[]int{1},
		//	[]int{1},
		//},
		//{
		//	[]int{},
		//	[]int{},
		//},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			head := createLinkedListFromSlice(tc.input)
			if !compareLinkedListValues(head, tc.input) {
				t.Fatalf("Your function that creates linked lists doesnt work")
			}

			got := reverseList(head)
			want := tc.want
			if !compareLinkedListValues(got, want) {
				t.Fatalf("got %v want %v", linkedListToSlice(got), want)
			}
		})
	}
}
