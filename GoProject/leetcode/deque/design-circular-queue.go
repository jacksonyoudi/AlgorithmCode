package deque

type MyCircularQueue struct {
	len   int
	data  []int
	front int
	rear  int
	size  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		len:   k,
		data:  make([]int, k),
		front: 0,
		rear:  0,
		size:  0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.data[this.rear] = value
	this.rear += 1
	if this.rear%this.len == 0 {
		this.rear = 0
	}
	this.size += 1
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.front += 1
	if this.front%this.len == 0 {
		this.front = 0
	}
	this.size -= 1
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.data[this.front]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	a :=    this.rear - 1
	if a < 0 {
		return this.data[this.len-1]
	}
	return this.data[a]
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if this.size == this.len {
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
