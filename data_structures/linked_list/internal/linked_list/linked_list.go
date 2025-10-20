package linked_list

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func (l *LinkedList[T]) AddFirst(val T) {
	newNode := Node[T]{val, l.head}
	l.head = &newNode
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

func (l *LinkedList[T]) RemoveFirst() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
}
