package main

type LinkNode struct {
	Next *LinkNode
	key  int
}

type MyHashSet struct {
	data []*LinkNode
	mod  int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	a := MyHashSet{}
	a.mod = 16
	a.data = make([]*LinkNode, a.mod)

	return a
}

func (this *MyHashSet) Add(key int) {
	index := key % this.mod
	node := this.data[index]
	if node != nil {
		for node != nil {
			if node.key == key {
				return
			}
			node = node.Next
		}
		newNode := new(LinkNode)
		newNode.key = key
		node.Next = newNode
	} else {
		node = new(LinkNode)
		node.key = key
		this.data[index] = node
	}

}

func (this *MyHashSet) Remove(key int) {
	index := key % this.mod
	node := this.data[index]
	if node != nil {
		preNode := node
		for node != nil {
			if node.key == key {
				if preNode != node {
					// 删除
					preNode.Next = node.Next
					return
				} else {
					this.data[index] = node.Next
					return
				}
			}
			preNode = node
			node = node.Next
		}
	}
	return

}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	index := key % this.mod
	node := this.data[index]
	if node != nil {
		for node != nil {
			if node.key == key {
				return true
			}
			node = node.Next
		}
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

func main() {

}
