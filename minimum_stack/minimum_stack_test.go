package minimum_stack

import "testing"

func TestMinStack(t *testing.T) {
	minStack := Constructor()

	minStack.Push(1)
	minStack.Push(2)
	minStack.Push(0)

	got := minStack.GetMin()
	want := 0
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	minStack.Pop()
	got = minStack.Top()
	want = 2
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	got = minStack.GetMin()
	want = 1
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMinStack2(t *testing.T) {
	minStack := Constructor()

	minStack.Push(10)
	minStack.Pop()
	minStack.Push(20)

	got := minStack.Top()
	want := 20
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	minStack.Push(-20)
	got = minStack.GetMin()
	want = -20
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
