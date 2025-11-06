package implement_stack_using_queues

import "testing"

func TestMyStack(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)

	got := stack.Top()
	want := 2
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	got = stack.Pop()
	want = 2
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	empty := stack.Empty()
	if empty != false {
		t.Fatalf("got %v, want %v", got, want)
	}

	got = stack.Pop()
	want = 1
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	empty = stack.Empty()
	if empty != true {
		t.Fatalf("got %v, want %v", got, want)
	}
}
