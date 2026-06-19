package copy_linked_list_with_random_pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	randomPointerMap := make(map[*Node]*Node)

	// form the deep copy linked list, with the random keys pointing to the original random values
	// on iteration, add the key value pair of: orig address|deep copy address of node into the hash map
	dummy := &Node{}
	savedHead := dummy

	for head != nil {
		deepCopyNode := &Node{
			Val:    head.Val,
			Next:   nil,
			Random: head.Random,
		}
		dummy.Next = deepCopyNode

		randomPointerMap[head] = deepCopyNode

		dummy = dummy.Next
		head = head.Next
	}

	// on 2nd pass, go through each node and replace the orig address with the deep copy address using the hash map
	dummy = savedHead
	dummy = dummy.Next
	for dummy != nil {
		dummy.Random = randomPointerMap[dummy.Random]
		dummy = dummy.Next
	}

	return savedHead.Next
}
