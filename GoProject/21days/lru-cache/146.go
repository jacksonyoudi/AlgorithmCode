package lru_cache

type Node struct {
	Key, Val int
	Next     *Node
	Prev     *Node
}

type DoubleList struct {
	// 头尾虚节点
	Head, Tail *Node
	// 链表元素数
	Size int
}

func NewDoubleList() *DoubleList {
	n := &DoubleList{
		Head: &Node{0, 0, nil, nil},
		Tail: &Node{0, 0, nil, nil},
		Size: 0,
	}

	n.Head.Next = n.Tail
	n.Tail.Prev = n.Head

	return n

}

// 在链表尾部添加节点 x，时间 O(1)
func (d *DoubleList) addLast(node *Node) {
	// 头尾都是虚节点， 这样处理就可以 不用多考虑了
	node.Prev = d.Tail.Prev
	node.Next = d.Tail
	d.Tail.Prev.Next = node
	d.Tail.Prev = node
	d.Size++

}

// 删除链表中的 x 节点（x 一定存在）
// 由于是双链表且给的是目标 Node 节点，时间 O(1)
func (d *DoubleList) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	d.Size--

}

// 删除链表中第一个节点，并返回该节点，时间 O(1)
func (d *DoubleList) removeFirst() *Node {
	// 判断是否为空
	if d.Head.Next == d.Tail {
		return nil
	}
	first := d.Head.Next
	d.remove(first)
	return first
}

type LRUCache struct {
	Cap   int
	Dmap  map[int]*Node
	Cache *DoubleList
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:   capacity,
		Dmap:  map[int]*Node{},
		Cache: NewDoubleList(),
	}
}

// 将某个 key 提升为最近使用的
// 一定存在
func (this *LRUCache) makeRecently(key int) {
	node, _ := this.Dmap[key]
	this.Cache.remove(node)
	this.Cache.addLast(node)
}

// 添加元素
func (this *LRUCache) addRecently(key, val int) {
	node := &Node{key, val, nil, nil}
	this.Cache.addLast(node)
	this.Dmap[key] = node
}

// 删除某一个 key
func (this *LRUCache) deleteKey(key int) {
	node, _ := this.Dmap[key]
	this.Cache.remove(node)
	delete(this.Dmap, key)
}

/* 删除最久未使用的元素 */
func (this *LRUCache) removeLeastRecently() {
	first := this.Cache.removeFirst()
	delete(this.Dmap, first.Key)
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Dmap[key]; ok {
		this.makeRecently(key)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 先判断是否在
	if node, ok := this.Dmap[key]; ok {
		this.makeRecently(key)
		node.Val = value
		return
	}
	if len(this.Dmap) == this.Cap {
		this.removeLeastRecently()
	}
	this.addRecently(key, value)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
