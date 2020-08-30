/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
type nodeLRU struct {
	key  int
	val  int
	prev *nodeLRU
	next *nodeLRU
}

type LRUCache struct {
	cap   int
	cache map[int]*nodeLRU
	head  *nodeLRU
	tail  *nodeLRU
}

func Constructor(capacity int) LRUCache {
	head := &nodeLRU{}
	tail := &nodeLRU{}

	head.next = tail
	tail.prev = head

	return LRUCache{
		cap:   capacity,
		cache: make(map[int]*nodeLRU, capacity),
		head:  head,
		tail:  tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.visit(key)
		return node.val
	}

	return -1
}

func (this *LRUCache) visit(key int) {
	node := this.cache[key]

	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = this.tail
	node.prev = this.tail.prev

	node.next.prev = node
	node.prev.next = node
}

func (this *LRUCache) evict() {
	node := this.head.next
	delete(this.cache, node.key)

	node.prev.next = node.next
	node.next.prev = this.head
}

func (this *LRUCache) Put(key int, value int) {
	if this.cap <= 0 {
		return
	}

	// exist, update
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.visit(key)
		return
	}

	// not exist, insert
	if len(this.cache) >= this.cap { // for <= can deal with lazy evict when this.cap change
		this.evict()
	}

	node := &nodeLRU{key: key, val: value}
	this.cache[key] = node

	node.next = this.tail
	node.prev = this.tail.prev

	node.next.prev = node
	node.prev.next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

