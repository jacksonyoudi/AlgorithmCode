package design_linked_list

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

type MyLinkedList struct {
	Length int
	Head   *Node
	Tail   *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		0,
		nil,
		nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index > this.Length-1 {
		return -1
	}

	i := 0
	cur := this.Head
	for i < index {
		cur = cur.Next
		i++
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	n := &Node{
		val,
		nil,
		nil,
	}
	h := this.Head

	if this.Length == 0 {
		this.Head = n
		this.Tail = n
	} else {
		this.Head = n
		this.Head.Next = h
		this.Head.Prev = nil
		h.Prev = n
	}
	this.Length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	n := &Node{
		val,
		nil,
		nil,
	}
	t := this.Tail

	if this.Length == 0 {
		this.Head = n
		this.Tail = n
	} else {
		t.Next = n
		this.Tail = n
		n.Prev = t
	}
	this.Length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	} else if index == this.Length {
		this.AddAtTail(val)
		return
	} else if index > this.Length {
		return
	} else {
		i := 0
		cur := this.Head
		for i < index {
			cur = cur.Next
			i++

		}

		n := &Node{
			val,
			nil,
			nil,
		}

		next := cur
		pre := cur.Prev

		pre.Next = n
		next.Prev = n
		n.Prev = pre
		n.Next = next
	}
	this.Length++

}

func (this *MyLinkedList) DeleteHead() {
	if this.Length == 0 {
		return
	}

	cur := this.Head
	next := cur.Next

	if this.Length == 1 {
		this.Tail = nil
		this.Head = nil
		this.Length--
		return
	}

	this.Head = next
	this.Head.Prev = nil
	this.Length--
}

func (this *MyLinkedList) DeleteTail() {
	if this.Length == 0 {
		return
	}
	cur := this.Tail
	prev := cur.Prev

	this.Tail = prev
	this.Tail.Next = nil
	this.Length--
	if this.Length == 1 {
		this.Tail = this.Head
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index > this.Length-1 {
		return
	} else if index == 0 {
		this.DeleteHead()
		return
	} else if index == this.Length-1 {
		this.DeleteTail()
		return
	}

	cur := this.Head
	i := 0
	for i < index {
		cur = cur.Next
		i++
	}

	next := cur.Next
	pre := cur.Prev

	pre.Next = next
	next.Prev = pre
	this.Length--
}
