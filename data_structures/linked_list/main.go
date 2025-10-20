package main

import "linked_list/internal/linked_list"

func main() {
	l := linked_list.LinkedList[int]{}

	l.AddFirst(3)
	l.AddFirst(2)
	l.AddFirst(1)

	l.Display()
}
