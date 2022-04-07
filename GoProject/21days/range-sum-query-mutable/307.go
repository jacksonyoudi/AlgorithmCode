package range_sum_query_mutable

type NumArray struct {
	nums   []int
	preSum []int
	size   int
}

func Constructor(nums []int) NumArray {
	size := len(nums)
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	sum := 0
	preSum[0] = sum
	for i, v := range nums {
		preSum[i+1] = v + sum
		sum = v + sum
	}

	return NumArray{
		nums:   nums,
		preSum: preSum,
		size:   size,
	}
}

func (this *NumArray) Update(index int, val int) {
	old := this.nums[index]
	this.nums[index] = val
	sub := val - old

	for i := index + 1; i <= this.size; i++ {
		this.preSum[i] = this.preSum[i] + sub
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
