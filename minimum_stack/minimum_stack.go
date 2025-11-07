package minimum_stack

type Node struct {
	data   int
	minVal int
}

type MinStack struct {
	s      []Node
	minVal int
}

func Constructor() MinStack {
	return MinStack{make([]Node, 0), 0}
}

func (this *MinStack) Push(val int) {
	n := Node{val, 0}
	if len(this.s) == 0 {
		this.minVal = val
	} else {
		this.minVal = min(val, this.minVal)
	}
	n.minVal = this.minVal
	this.s = append(this.s, n)
}

func (this *MinStack) Pop() {
	this.s = this.s[:len(this.s)-1]
	if len(this.s) != 0 {
		this.minVal = this.s[len(this.s)-1].minVal
	}
}

func (this *MinStack) Top() int {
	n := this.s[len(this.s)-1]
	return n.data
}

func (this *MinStack) GetMin() int {
	return this.minVal
}
