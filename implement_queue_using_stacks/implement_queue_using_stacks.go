package implement_queue_using_stacks

type MyQueue struct {
	s1 []int // stack to push to
	s2 []int // stack to pop from
}

func Constructor() MyQueue {
	return MyQueue{make([]int, 0), make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.s1 = append(this.s1, x)
}

func (this *MyQueue) Pop() int {
	if len(this.s2) == 0 {
		for range len(this.s1) {
			this.s2 = append(this.s2, this.s1[len(this.s1)-1])
			this.s1 = this.s1[:len(this.s1)-1]
		}
	}

	val := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return val
}

func (this *MyQueue) Peek() int {
	if len(this.s2) == 0 {
		for range len(this.s1) {
			this.s2 = append(this.s2, this.s1[len(this.s1)-1])
			this.s1 = this.s1[:len(this.s1)-1]
		}
	}

	val := this.s2[len(this.s2)-1]
	return val
}

func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0 && len(this.s2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
