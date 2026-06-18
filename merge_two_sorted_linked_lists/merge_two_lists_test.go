package merge_two_lists

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

func TestMergeTwoLists(t *testing.T) {
	tt := []struct {
		list1 []int
		list2 []int
		want  []int
	}{
		{
			[]int{1, 3, 5},
			[]int{2, 4, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{},
			[]int{1, 2},
			[]int{1, 2},
		},
	}

	for _, tc := range tt {
		t.Run("testcase", func(t *testing.T) {
			ll1 := createLinkedListFromSlice(tc.list1)
			ll2 := createLinkedListFromSlice(tc.list2)

			got := mergeTwoLists(ll1, ll2)
			want := tc.want

			if !compareLinkedListValues(got, want) {
				t.Fatalf("got %v want %v", linkedListToSlice(got), want)
			}
		})
	}
}
