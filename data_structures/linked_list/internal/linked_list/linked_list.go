package linked_list

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{}
}

// for debugging purposes
func (l *LinkedList[T]) Display() {
	tmp := l.head
	if tmp == nil {
		fmt.Println("<empty>")
		return
	}

	fmt.Print("head: ")
	for tmp != nil {
		fmt.Printf("%v", tmp.data)
		if tmp.next != nil {
			fmt.Print(" -> ")
		}
		tmp = tmp.next
	}
	fmt.Println()
}

func (l *LinkedList[T]) GetHead() T {
	return l.head.data
}

func (l *LinkedList[T]) GetTail() T {
	return l.tail.data
}

func (l *LinkedList[T]) GetLength() int {
	return l.length
}

func (l *LinkedList[T]) AddFirst(val T) {
	n := Node[T]{val, nil}

	if l.length == 0 {
		l.head = &n
		l.tail = &n
	} else {
		n.next = l.head
		l.head = &n
	}

	l.length++
}

func (l *LinkedList[T]) AddLast(val T) {
	n := Node[T]{val, nil}

	if l.length == 0 {
		l.head = &n
		l.tail = &n
	} else {
		l.tail.next = &n
		l.tail = &n
	}

	l.length++
}

func (l *LinkedList[T]) RemoveFirst() {
	if l.length == 0 {
		return
	}

	if l.length == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
	}

	l.length--
}

func (l *LinkedList[T]) RemoveLast() {
	if l.length == 0 {
		return
	}

	if l.length == 1 {
		l.head = nil
		l.tail = nil
	} else {
		tmp := l.head
		for tmp != nil {
			if tmp.next == l.tail {
				break
			}
			tmp = tmp.next
		}

		tmp.next = nil
		l.tail = tmp
	}

	l.length--
}

func (l *LinkedList[T]) Get(index int) (T, bool) {
	tmp := l.head
	i := 0
	for i < l.length || tmp != nil {
		if i == index {
			return tmp.data, true
		}
		i++
		tmp = tmp.next
	}

	return *new(T), false
}

func (l *LinkedList[T]) Set(index int, value T) bool {
	tmp := l.head
	i := 0
	for i < l.length || tmp != nil {
		if i == index {
			tmp.data = value
			return true
		}
		i++
		tmp = tmp.next
	}

	return false
}

func (l *LinkedList[T]) Insert(index int, value T) bool {
	// bounds check
	if index < 0 || index > l.length {
		return false
	}

	// insert into linked list start
	if index == 0 {
		l.AddFirst(value)
		return true
	}

	// insert into linked list end
	if index == l.length {
		l.AddLast(value)
		return true
	}

	// inserting into middle of linked list
	tmp := l.head
	for i := 0; i < index; i++ {
		tmp = tmp.next
	}
	n := Node[T]{value, tmp.next}
	tmp.next = &n
	l.length++

	return true
}

func (l *LinkedList[T]) Remove(index int) bool {
	// bounds check
	if index < 0 || index > l.length {
		return false
	}

	// remove first node
	if index == 0 {
		l.RemoveFirst()
		return true
	}

	// remove last node
	if index == l.length-1 {
		l.RemoveLast()
		return true
	}

	// remove node in middle of linked list
	tmp := l.head
	for i := 0; i < index-1; i++ {
		tmp = tmp.next
	}
	tmp.next = tmp.next.next
	l.length--

	return true
}

func (l *LinkedList[T]) Reverse() {
	if l.head == nil {
		return
	}

	var prev *Node[T]
	prev = nil
	tmp := l.head
	next := tmp.next

	for tmp != nil {
		tmp.next = prev
		prev = tmp
		tmp = next
		if next != nil {
			next = next.next
		}
	}

	tmp = l.head
	l.head = l.tail
	l.tail = l.head
}

func (l *LinkedList[T]) ConvertToSlice() []T {
	values := make([]T, 0, l.length)

	tmp := l.head
	for tmp != nil {
		values = append(values, tmp.data)
		tmp = tmp.next
	}

	return values
}
