package add_two_numbers

import "testing"

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

func TestRemoveNthFromEnd(t *testing.T) {
	tt := []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{5, 7, 9},
		},
		{
			[]int{9},
			[]int{9},
			[]int{8, 1},
		},
		{
			[]int{9, 1},
			[]int{9},
			[]int{8, 2},
		},
		{
			[]int{9, 1, 1},
			[]int{9},
			[]int{8, 2, 1},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			ll1 := createLinkedListFromSlice(tc.l1)
			ll2 := createLinkedListFromSlice(tc.l2)

			got := addTwoNumbers(ll1, ll2)
			want := tc.want

			if !compareLinkedListValues(got, want) {
				t.Fatalf("got %v want %v", linkedListToSlice(got), want)
			}
		})
	}
}
