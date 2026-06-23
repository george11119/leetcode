package kth_largest_integer

import (
	"container/heap"
)

type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] < h[j] } // ">" makes it a MAX heap. "<" will make it a MIN heap
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type KthLargest struct {
	h *IntMaxHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	h := make(IntMaxHeap, len(nums))
	copy(h, nums)
	heap.Init(&h)

	for i := 0; i < len(nums)-k; i++ {
		heap.Pop(&h)
	}

	return KthLargest{h: &h, k: k}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.h, val)
	if this.h.Len() > this.k {
		heap.Pop(this.h)
	}

	return (*this.h)[0]
}
