package design_circular_queue

type Node struct {
	Val  int
	Next *Node
}

type MyCircularQueue struct {
	maxSize int
	size    int
	head    *Node
	tail    *Node
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		maxSize: k,
		size:    0,
		head:    nil,
		tail:    nil,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	newNode := &Node{Val: value, Next: nil}

	if this.IsEmpty() {
		this.head = newNode
		this.tail = newNode
	} else {
		this.tail.Next = newNode
		this.tail = newNode
	}

	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.head = this.head.Next
	this.size--
	if this.IsEmpty() {
		this.tail = nil
	}

	return true
}

func (this *MyCircularQueue) Front() int {
	if this.head == nil {
		return -1
	}

	return this.head.Val
}

func (this *MyCircularQueue) Rear() int {
	if this.tail == nil {
		return -1
	}

	return this.tail.Val
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.maxSize
}
