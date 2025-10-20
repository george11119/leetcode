package linked_list

import "testing"

func TestAddFirst(t *testing.T) {
	l := LinkedList[int]{}
	l.AddFirst(1)
	l.AddFirst(2)

	got := l.head.data
	want := 2

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = l.head.next.data
	want = 1

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestRemoveFirst(t *testing.T) {
	l := LinkedList[int]{}
	l.AddFirst(1)
	l.AddFirst(2)
	// head: 2 -> 1

	got := l.head.data
	want := 2

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	l.RemoveFirst()
	// head: 1

	got = l.head.data
	want = 1

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	l.RemoveFirst()
	// head: nil

	headGot := l.head

	if headGot != nil {
		t.Fatalf("empty head is not nil. got %v", headGot)
	}
}
