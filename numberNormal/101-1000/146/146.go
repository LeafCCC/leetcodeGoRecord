package leetcode

type DualListNode struct {
	Val       int
	Pre, Next *DualListNode
}

type LRUCache struct {
	capacity int
	now      int
	record   map[int]*DualListNode
	toKey    map[*DualListNode]int
	head     *DualListNode
	tail     *DualListNode
}

func Constructor(capacity int) LRUCache {
	res := LRUCache{capacity: capacity, now: 0}
	res.record = make(map[int]*DualListNode)
	res.toKey = make(map[*DualListNode]int)
	head, tail := &DualListNode{}, &DualListNode{}
	head.Next = tail
	tail.Pre = head
	res.head, res.tail = head, tail
	return res
}

func (this *LRUCache) Get(key int) int {
	if this.record[key] == nil {
		return -1
	}

	where := this.record[key]
	where.Pre.Next, where.Next.Pre = where.Next, where.Pre
	this.head.Next, this.head.Next.Pre, where.Pre, where.Next = where, where, this.head, this.head.Next
	return where.Val
}

func (this *LRUCache) Put(key int, value int) {
	if this.record[key] != nil {
		where := this.record[key]
		where.Pre.Next, where.Next.Pre = where.Next, where.Pre
		this.now--
	} else if this.now == this.capacity {
		tmpKey := this.toKey[this.tail.Pre]
		this.record[tmpKey] = nil
		this.toKey[this.tail.Pre] = -1
		this.tail.Pre.Pre.Next, this.tail.Pre = this.tail, this.tail.Pre.Pre
		this.now--
	}

	newCache := &DualListNode{Val: value}
	this.toKey[newCache] = key
	this.record[key] = newCache
	head := this.head
	head.Next, head.Next.Pre, newCache.Pre, newCache.Next = newCache, newCache, head, head.Next
	this.now++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
