package maximum_frequency_stack

type FreqStack struct {
	max   int
	count map[int]int
	stack map[int][]int
}

func Constructor() FreqStack {
	return FreqStack{0, make(map[int]int), make(map[int][]int)}
}

func (this *FreqStack) Push(val int) {
	this.count[val]++
	if this.count[val] > this.max {
		this.max = this.count[val]
	}
	this.stack[this.count[val]] = append(this.stack[this.count[val]], val)
}

func (this *FreqStack) Pop() int {
	res := this.stack[this.max][len(this.stack[this.max])-1]
	this.stack[this.max] = this.stack[this.max][:len(this.stack[this.max])-1]
	this.count[res]--
	if len(this.stack[this.max]) == 0 {
		this.max--
	}
	return res
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor()
 * obj.Push(val)
 * param2 := obj.Pop()
 */
