package lru_cache

type ListNode struct {
	key  int
	data int
	next *ListNode
	prev *ListNode
}

type LRUCache struct {
	capacity int
	cache    map[int]*ListNode
	head     *ListNode
	tail     *ListNode
}

func (this *LRUCache) InsertTail(newNode *ListNode) {
	this.tail.prev.next = newNode
	newNode.prev = this.tail.prev
	newNode.next = this.tail
	this.tail.prev = newNode
}

func (this *LRUCache) RemoveNode(node *ListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func Constructor(capacity int) LRUCache {
	headNode := &ListNode{}
	tailNode := &ListNode{}
	headNode.next = tailNode
	tailNode.prev = headNode

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*ListNode),
		head:     headNode,
		tail:     tailNode,
	}
}

func (this *LRUCache) Get(key int) int {
	node, found := this.cache[key]
	if !found {
		return -1
	}

	this.RemoveNode(node)
	this.InsertTail(node)
	return node.data
}

func (this *LRUCache) Put(key int, value int) {
	node, found := this.cache[key]
	if found {
		this.RemoveNode(node)
		delete(this.cache, node.key)
	}

	newNode := &ListNode{key: key, data: value}
	this.InsertTail(newNode)
	this.cache[key] = newNode

	if len(this.cache) > this.capacity {
		lru := this.head.next
		this.RemoveNode(lru)
		delete(this.cache, lru.key)
	}
}
