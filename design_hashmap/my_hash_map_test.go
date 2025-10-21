package my_hash_map

import "testing"

func assert(t *testing.T, got, want any) {
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestMyHashMap(t *testing.T) {
	m := Constructor()
	m.Put(1, 1)
	m.Put(2, 2)

	assert(t, m.Get(1), 1)
	assert(t, m.Get(3), -1)

	m.Put(2, 1)
	assert(t, m.Get(2), 1)

	m.Remove(2)
	assert(t, m.Get(2), -1)
}
