package reorder_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	// get list length
	llLen := 0
	tmp := head
	for tmp != nil {
		llLen++
		tmp = tmp.Next
	}

	// get middle node
	tmp = head
	for i := 0; i < llLen/2; i++ {
		tmp = tmp.Next
	}

	// reverse the 2nd half of the list
	var prev *ListNode
	curr := tmp
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	revHead := prev

	dummy := &ListNode{}
	tail := dummy

	for revHead != nil && head != nil {
		if head == revHead {
			break
		}
		tail.Next = head
		head = head.Next
		tail = tail.Next

		tail.Next = revHead
		revHead = revHead.Next
		tail = tail.Next
	}
}
