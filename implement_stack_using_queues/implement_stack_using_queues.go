package implement_stack_using_queues

// a queue
type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{make([]T, 0)}
}

func (this *Queue[T]) Size() int {
	return len(this.data)
}

func (this *Queue[T]) IsEmpty() bool {
	return this.Size() == 0
}

// push to back of queue
func (this *Queue[T]) Push(x T) {
	this.data = append(this.data, x)
}

func (this *Queue[T]) Peek() T {
	return this.data[0]
}

func (this *Queue[T]) Pop() T {
	val := this.data[0]
	this.data = this.data[1:]
	return val
}

type MyStack struct {
	data Queue[int]
}

func Constructor() MyStack {
	return MyStack{NewQueue[int]()}
}

func (this *MyStack) Push(x int) {
	this.data.Push(x)
}

func (this *MyStack) Pop() int {
	for i := 0; i < this.data.Size()-1; i++ {
		val := this.data.Pop()
		this.data.Push(val)
	}
	res := this.data.Pop()
	return res
}

func (this *MyStack) Top() int {
	for i := 0; i < this.data.Size()-1; i++ {
		val := this.data.Pop()
		this.data.Push(val)
	}
	res := this.data.Pop()
	this.data.Push(res)
	return res
}

func (this *MyStack) Empty() bool {
	return this.data.IsEmpty()
}
