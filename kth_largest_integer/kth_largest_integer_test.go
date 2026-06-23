package kth_largest_integer

import (
	"testing"
)

//func TestExample(t *testing.T) {
//	h := &IntMaxHeap{3, 1, 4, 1, 5, 9, 2, 6}
//	heap.Init(h)
//
//	heap.Push(h, 8)
//
//	for h.Len() > 0 {
//		fmt.Printf("%d ", heap.Pop(h)) // 9 8 6 5 4 3 2 1 1
//	}
//}

func TestKthLargest(t *testing.T) {
	kthLargest := Constructor(3, []int{1, 2, 3, 3})

	var got int
	var want int

	got = kthLargest.Add(3)
	want = 3
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = kthLargest.Add(5)
	want = 3
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = kthLargest.Add(6)
	want = 3
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = kthLargest.Add(7)
	want = 5
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = kthLargest.Add(8)
	want = 6
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestKthLargest2(t *testing.T) {
	kthLargest := Constructor(5, []int{-1, -1, -1, -1, -1, -1})

	var got int
	var want int

	got = kthLargest.Add(-1)
	want = -1
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = kthLargest.Add(-1)
	want = -1
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = kthLargest.Add(-1)
	want = -1
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}
