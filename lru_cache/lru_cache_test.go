package lru_cache

import (
	"testing"
)

func assert(t *testing.T, got any, want any) {
	t.Helper()

	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestLRUCache(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 10)
	assert(t, lruCache.Get(1), 10)
	lruCache.Put(2, 20)
	lruCache.Put(3, 30)
	assert(t, lruCache.Get(2), 20)
	assert(t, lruCache.Get(1), -1)
}

func TestLRUCache2(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	assert(t, lruCache.Get(1), 1)
	lruCache.Put(3, 3)
	assert(t, lruCache.Get(2), -1)
	lruCache.Put(4, 4)
	assert(t, lruCache.Get(1), -1)
	assert(t, lruCache.Get(3), 3)
	assert(t, lruCache.Get(4), 4)
}

func TestLRUCache3(t *testing.T) {
	lruCache := Constructor(3)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)
	assert(t, lruCache.Get(1), 1)
	assert(t, lruCache.Get(2), 2)
	assert(t, lruCache.Get(4), -1)
	lruCache.Put(4, 4)
	assert(t, lruCache.Get(1), 1)
	assert(t, lruCache.Get(2), 2)
	assert(t, lruCache.Get(3), -1)
	assert(t, lruCache.Get(4), 4)
	assert(t, lruCache.Get(2), 2)
	lruCache.Put(1, 8)
	lruCache.Put(3, 7)
	assert(t, lruCache.Get(1), 8)
	assert(t, lruCache.Get(2), 2)
	assert(t, lruCache.Get(3), 7)
	assert(t, lruCache.Get(4), -1)
	assert(t, lruCache.Get(5), -1)
	assert(t, lruCache.Get(2), 2)
	assert(t, lruCache.Get(3), 7)
	assert(t, lruCache.Get(4), -1)
	lruCache.Put(1, 9)
	lruCache.Put(6, 6)
	assert(t, lruCache.Get(1), 9)
	assert(t, lruCache.Get(2), -1)
	assert(t, lruCache.Get(3), 7)
	assert(t, lruCache.Get(4), -1)
	assert(t, lruCache.Get(5), -1)
	assert(t, lruCache.Get(6), 6)
}
