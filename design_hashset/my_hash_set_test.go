package my_hash_set

import (
	"testing"
)

func assert(t *testing.T, got, want any) {
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestHashSet(t *testing.T) {
	h := Constructor()
	h.Add(1)
	h.Add(2)

	ok := h.Contains(1)
	assert(t, ok, true)

	ok = h.Contains(3)
	assert(t, ok, false)

	h.Add(2)

	ok = h.Contains(2)
	assert(t, ok, true)

	h.Remove(2)

	ok = h.Contains(2)
	assert(t, ok, false)
}
