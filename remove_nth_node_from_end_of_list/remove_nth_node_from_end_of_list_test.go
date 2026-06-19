package remove_nth_node_from_end_of_list

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
		l    []int
		n    int
		want []int
	}{
		{
			[]int{1, 2, 3, 4},
			2,
			[]int{1, 2, 4},
		},
		{
			[]int{5},
			1,
			[]int{},
		},
		{
			[]int{1, 2},
			2,
			[]int{2},
		},
		{
			[]int{1, 2},
			1,
			[]int{1},
		},
		{
			[]int{1, 2, 3},
			3,
			[]int{2, 3},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			ll := createLinkedListFromSlice(tc.l)
			got := removeNthFromEnd(ll, tc.n)
			want := tc.want

			if !compareLinkedListValues(got, want) {
				t.Fatalf("got %v want %v", linkedListToSlice(got), want)
			}
		})
	}
}
