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
