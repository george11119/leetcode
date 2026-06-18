package linked_list_cycle_detection

import "testing"

func TestHasCycle(t *testing.T) {
	l := &ListNode{1, nil}
	cycleNode := &ListNode{2, nil}
	l.Next = cycleNode
	l = l.Next
	l.Next = &ListNode{3, nil}
	l = l.Next
	l.Next = &ListNode{4, nil}
	l = l.Next
	l.Next = cycleNode

	got := hasCycle(l)
	want := true

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}
