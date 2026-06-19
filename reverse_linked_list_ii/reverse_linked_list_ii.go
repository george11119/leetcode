package reverse_linked_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode, count int) *ListNode {
	var prev *ListNode
	var savedHead *ListNode
	var savedNext *ListNode
	for i := 0; i < count; i++ {
		if i == 0 {
			savedHead = head
		}
		if i == count-1 {
			savedNext = head.Next
		}
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	savedHead.Next = savedNext
	return prev
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	tmp := dummy
	for i := 0; i < left-1; i++ {
		tmp = tmp.Next
	}
	tmp.Next = reverse(tmp.Next, right-left+1)

	return dummy.Next
}
