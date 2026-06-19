package design_circular_queue

import (
	"testing"
)

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param1 := obj.EnQueue(value);
 * param2 := obj.DeQueue();
 * param3 := obj.Front();
 * param4 := obj.Rear();
 * param5 := obj.IsEmpty();
 * param6 := obj.IsFull();
 */

func assert(t *testing.T, got any, want any) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestCircularQueue(t *testing.T) {
	q := Constructor(3)

	assert(t, q.EnQueue(1), true)
	assert(t, q.EnQueue(2), true)
	assert(t, q.EnQueue(3), true)
	assert(t, q.EnQueue(4), false)

	assert(t, q.Rear(), 3)
	assert(t, q.IsFull(), true)
	assert(t, q.DeQueue(), true)
	assert(t, q.EnQueue(4), true)
	assert(t, q.Rear(), 4)
}

func TestCircularQueue2(t *testing.T) {
	q := Constructor(8)

	assert(t, q.EnQueue(3), true)
	assert(t, q.EnQueue(9), true)
	assert(t, q.EnQueue(5), true)
	assert(t, q.EnQueue(0), true)

	assert(t, q.DeQueue(), true)
	assert(t, q.DeQueue(), true)
	assert(t, q.IsEmpty(), false)
	assert(t, q.IsEmpty(), false)
	assert(t, q.Rear(), 0)
	assert(t, q.Rear(), 0)
	assert(t, q.DeQueue(), true)
}
