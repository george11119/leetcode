package linked_list

import "testing"

func TestAddFirst(t *testing.T) {
	l := NewLinkedList[int]()
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

	got = l.tail.data
	want = 1

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = l.GetLength()
	want = 2

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestAddLast(t *testing.T) {
	l := NewLinkedList[int]()
	l.AddLast(1)
	l.AddLast(2)

	got := l.head.data
	want := 1

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = l.head.next.data
	want = 2

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = l.tail.data
	want = 2

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = l.GetLength()
	want = 2

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestRemoveFirst(t *testing.T) {
	l := NewLinkedList[int]()
	l.AddFirst(1)
	l.AddFirst(2)
	// head: 2 -> 1

	got := l.head.data
	want := 2

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = l.GetLength()
	want = 2
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

	got = l.GetLength()
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

	got = l.GetLength()
	want = 0
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	l.RemoveFirst()

	got = l.GetLength()
	want = 0
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestRemoveLast(t *testing.T) {
	l := NewLinkedList[int]()

	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)

	// head: 1 -> 2 -> 3

	got := l.tail.data
	want := 3

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = l.GetLength()
	want = 3
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	l.RemoveLast()
	// head: 1 -> 2

	got = l.head.data
	want = 2

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = l.GetLength()
	want = 2
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	// l.RemoveLast()
	// // head: 1

	// got = l.head.data
	// want = 1

	// if got != want {
	// 	t.Fatalf("got %v want %v", got, want)
	// }

	// got = l.GetLength()
	// want = 1
	// if got != want {
	// 	t.Fatalf("got %v want %v", got, want)
	// }

	// l.RemoveLast()
	// // head: nil

	// headGot := l.head

	// if headGot != nil {
	// 	t.Fatalf("empty head is not nil. got %v", headGot)
	// }

	// got = l.GetLength()
	// want = 0
	// if got != want {
	// 	t.Fatalf("got %v want %v", got, want)
	// }

	// l.RemoveFirst()

	// got = l.GetLength()
	// want = 0
	// if got != want {
	// 	t.Fatalf("got %v want %v", got, want)
	// }
}
