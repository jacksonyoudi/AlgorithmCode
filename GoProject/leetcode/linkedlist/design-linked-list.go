package linkedlist

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	dummy *Node
	size  int
	tail  *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	dummpy := Node{}
	return MyLinkedList{
		dummy: &dummpy,
		size:  0,
		tail:  &dummpy,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if this.size <= index || index < 0 {
		return -1
	}

	return this.GetPreNode(index).next.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{val, nil}
	node.next = this.dummy.next
	this.dummy.next = node
	if this.size == 0 {
		this.tail = node
	}
	this.size += 1
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{val, nil}
	this.tail.next = node
	this.tail = node
	this.size += 1

}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	node := &Node{val, nil}
	if index > this.size {
		return
	} else if index == this.size {
		this.AddAtTail(val)
	} else if index <= 0 {
		this.AddAtHead(val)
	} else {
		pre := this.GetPreNode(index)
		node.next = pre.next
		pre.next = node
		this.size += 1
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.size <= index || index < 0 {
		return
	}
	pre := this.GetPreNode(index)
	if this.tail == pre.next {
		this.tail = pre
	}

	pre.next = pre.next.next
	this.size -= 1
}

func (this *MyLinkedList) GetPreNode(index int) *Node {

	front := this.dummy.next
	back := this.dummy

	for i := 0; i < index; i++ {
		back = front
		front = front.next
	}

	return back

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
