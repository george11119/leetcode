package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0

	dummy := &ListNode{}
	cur := dummy

	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry

		cur.Next = &ListNode{Val: sum % 10, Next: nil}

		if sum/10 != 0 {
			carry = 1
		} else {
			carry = 0
		}

		l1 = l1.Next
		l2 = l2.Next
		cur = cur.Next
	}

	for l1 != nil {
		sum := l1.Val + carry

		cur.Next = &ListNode{Val: sum % 10, Next: nil}

		if sum/10 != 0 {
			carry = 1
		} else {
			carry = 0
		}

		l1 = l1.Next
		cur = cur.Next
	}

	for l2 != nil {
		sum := l2.Val + carry

		cur.Next = &ListNode{Val: sum % 10, Next: nil}

		if sum/10 != 0 {
			carry = 1
		} else {
			carry = 0
		}

		l2 = l2.Next
		cur = cur.Next
	}

	if carry == 1 {
		cur.Next = &ListNode{Val: 1, Next: nil}
	}

	return dummy.Next
}
