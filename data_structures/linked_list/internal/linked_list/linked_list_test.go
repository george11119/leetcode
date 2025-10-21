package linked_list

import (
	"reflect"
	"testing"
)

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

	l.RemoveLast()
	// head: 1

	got = l.tail.data
	want = 1

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got = l.GetLength()
	want = 1
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	l.RemoveLast()
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

func TestGet(t *testing.T) {
	l := NewLinkedList[int]()
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)
	l.AddLast(4)

	got, ok := l.Get(0)
	want := 1

	if !ok {
		t.Fatalf("get should have returned true for ok, got ok: %v", ok)
	}

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	got, ok = l.Get(3)
	want = 4

	if !ok {
		t.Fatalf("get should have returned true for ok, got ok: %v", ok)
	}

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	_, ok = l.Get(4)
	expected_ok := false

	if ok != expected_ok {
		t.Fatalf("got %v want %v", ok, expected_ok)
	}
}

func TestSet(t *testing.T) {
	l := NewLinkedList[int]()
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)
	l.AddLast(4)

	ok := l.Set(0, 10)
	got, _ := l.Get(0)
	want := 10

	if !ok {
		t.Fatalf("get should have returned true for ok, got ok: %v", ok)
	}

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	ok = l.Set(3, 42)
	got, _ = l.Get(3)
	want = 42

	if !ok {
		t.Fatalf("get should have returned true for ok, got ok: %v", ok)
	}

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}

	ok = l.Set(4, 100)
	expected_ok := false

	if ok != expected_ok {
		t.Fatalf("got %v want %v", ok, expected_ok)
	}

	ok = l.Set(-1, 100)
	expected_ok = false

	if ok != expected_ok {
		t.Fatalf("got %v want %v", ok, expected_ok)
	}
}

func TestInsert(t *testing.T) {
	l := NewLinkedList[int]()
	l.Insert(0, 1)
	l.Insert(1, 2)
	l.Insert(2, 3)

	got := l.ConvertToSlice()
	want := []int{1, 2, 3}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	ok := l.Insert(4, 5)
	if ok {
		t.Fatalf("wanted ok to be false. got %v", ok)
	}

	ok = l.Insert(-1, 5)
	if ok {
		t.Fatalf("wanted ok to be false. got %v", ok)
	}

	l.Insert(1, 5)

	got = l.ConvertToSlice()
	want = []int{1, 2, 5, 3}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	if l.length != 4 {
		t.Fatalf("length is wrong. got %v want 4", l.length)
	}
}

func TestRemove(t *testing.T) {
	l := NewLinkedList[int]()
	ok := l.Remove(-1)
	if ok {
		t.Fatalf("wanted ok to be false. got %v", ok)
	}

	ok = l.Remove(1)
	if ok {
		t.Fatalf("wanted ok to be false. got %v", ok)
	}

	l.AddLast(1)
	l.Remove(0)

	got := l.ConvertToSlice()
	want := []int{}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	l.AddLast(1)
	l.AddLast(2)
	l.Remove(1)

	got = l.ConvertToSlice()
	want = []int{1}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	l.AddLast(2)
	l.AddLast(3)
	l.Remove(1)

	got = l.ConvertToSlice()
	want = []int{1, 3}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	if l.length != 2 {
		t.Fatalf("length is wrong. got %v want 2", l.length)
	}
}

func TestReverse(t *testing.T) {
	l := NewLinkedList[int]()
	l.Reverse()

	l.AddLast(1)
	got := l.ConvertToSlice()
	want := []int{1}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	l.AddLast(2)
	l.AddLast(3)

	l.Reverse()
	got = l.ConvertToSlice()
	want = []int{3, 2, 1}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestConvertToSlice(t *testing.T) {
	l := NewLinkedList[int]()
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)

	got := l.ConvertToSlice()
	want := []int{1, 2, 3}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
