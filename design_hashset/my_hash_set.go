package my_hash_set

type Node struct {
	val  int
	next *Node
}

type MyHashSet struct {
	set []Node
}

func Constructor() MyHashSet {
	return MyHashSet{make([]Node, 10_000)}
}

func (this *MyHashSet) Add(key int) {
	tmp := &this.set[key%10000]

	for tmp.next != nil {
		if tmp.next.val == key {
			return
		}
		tmp = tmp.next
	}

	tmp.next = &Node{key, nil}
}

func (this *MyHashSet) Remove(key int) {
	tmp := &this.set[key%10000]
	for tmp.next != nil {
		if tmp.next.val == key {
			tmp.next = tmp.next.next
			return
		}
		tmp = tmp.next
	}
}

func (this *MyHashSet) Contains(key int) bool {
	tmp := &this.set[key%10000]
	for tmp.next != nil {
		if tmp.next.val == key {
			return true
		}
		tmp = tmp.next
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
