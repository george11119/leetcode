package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var cur, prev, next *ListNode
	cur = head
	next = nil
	prev = nil

	if cur == nil {
		return nil
	}

	for {
		next = cur.Next
		cur.Next = prev
		if next == nil {
			break
		}
		prev = cur
		cur = next
	}

	return cur
}
