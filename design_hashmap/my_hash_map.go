package my_hash_map

type Node struct {
	key  int
	val  int
	next *Node
}

type MyHashMap struct {
	m []Node
}

func Constructor() MyHashMap {
	return MyHashMap{make([]Node, 10000)}
}

func (this *MyHashMap) Put(key int, value int) {
	tmp := &this.m[key%10000]
	for tmp.next != nil {
		if tmp.next.key == key {
			tmp.next.val = value
			return
		}
		tmp = tmp.next
	}
	tmp.next = &Node{key, value, nil}
}

func (this *MyHashMap) Get(key int) int {
	tmp := &this.m[key%10000]
	for tmp.next != nil {
		if tmp.next.key == key {
			return tmp.next.val
		}
		tmp = tmp.next
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	tmp := &this.m[key%10000]
	for tmp.next != nil {
		if tmp.next.key == key {
			tmp.next = tmp.next.next
			return
		}
		tmp = tmp.next
	}
}
