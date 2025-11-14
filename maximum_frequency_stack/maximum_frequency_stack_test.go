package maximum_frequency_stack

import (
	"fmt"
	"testing"
)

func TestMaxFreqStack(t *testing.T) {
	f := Constructor()
	f.Push(5)
	f.Push(7)
	f.Push(5)
	f.Push(7)
	f.Push(4)
	f.Push(5)

	got := f.Pop()
	want := 5
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = f.Pop()
	want = 7
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = f.Pop()
	want = 5
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = f.Pop()
	want = 4
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}
func TestMaxFreqStack2(t *testing.T) {
	f := Constructor()
	f.Push(1)
	f.Push(1)
	f.Push(1)
	f.Push(2)

	got := f.Pop()
	want := 1
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = f.Pop()
	want = 1
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	f.Push(2)
	f.Push(2)
	f.Push(1)

	fmt.Println(f)

	got = f.Pop()
	want = 2
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = f.Pop()
	want = 1
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = f.Pop()
	want = 2
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}
