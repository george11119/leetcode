package implement_queue_using_stacks

import "testing"

func TestMyQueue(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)

	frontVal := q.Peek()
	if frontVal != 1 {
		t.Fatalf("got %v want %v", frontVal, 1)
	}

	frontVal = q.Pop()
	if frontVal != 1 {
		t.Fatalf("got %v want %v", frontVal, 1)
	}

	qEmpty := q.Empty()
	if qEmpty != false {
		t.Fatalf("got %v want %v", qEmpty, false)
	}

	frontVal = q.Peek()
	if frontVal != 2 {
		t.Fatalf("got %v want %v", frontVal, 1)
	}

	frontVal = q.Pop()
	if frontVal != 2 {
		t.Fatalf("got %v want %v", frontVal, 1)
	}

	qEmpty = q.Empty()
	if qEmpty != true {
		t.Fatalf("got %v want %v", qEmpty, true)
	}
}
