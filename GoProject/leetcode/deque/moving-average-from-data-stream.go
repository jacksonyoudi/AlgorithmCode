package deque

type MovingAverage struct {
	size  int
	len   int
	data  []int
	front int
	rear  int
	sum   int
}

/** Initialize your data structure here. */
func ConstructorA(size int) MovingAverage {
	return MovingAverage{
		size:  0,
		data:  make([]int, size),
		front: 0,
		rear:  0,
		len:   size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.sum += val
	if this.size != this.len {
		this.data[this.rear] = val
		this.rear += 1
		if this.rear == this.len {
			this.rear = 0
		}
		this.size += 1
	} else {
		a := this.data[this.front]
		this.sum -= a
		this.front += 1
		if this.front == this.len {
			this.front = 0
		}

		this.data[this.rear] = val
		this.rear += 1
		if this.rear == this.len {
			this.rear = 0
		}


	}

	return float64(this.sum) / float64(this.size)

}
