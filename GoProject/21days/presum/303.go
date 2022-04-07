package presum

type NumArray struct {
	presum []int
	nums   []int
}

func Constructor(nums []int) NumArray {
	presum := make([]int, 0)
	acc := 0
	presum = append(presum, acc)
	for i := 0; i < len(nums); i++ {
		acc = acc + nums[i]
		presum = append(presum, acc)
	}
	return NumArray{
		presum: presum,
		nums:   nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.presum[right+1] - this.presum[left]
}
