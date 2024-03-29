package _200

type PermuteData struct {
	data [][]int
}

func permuteBfs(d *PermuteData, tmpR []int, nums []int) {
	if len(nums) == 0 {
		d.data = append(d.data, tmpR)
		return

	}

	for i, v := range nums {
		tR := make([]int, 0)
		tR = append(tR,  tmpR...)

		tmp := make([]int, 0)
		tmp = append(tmp, nums[0:i]...)
		tmp = append(tmp, nums[i+1:]...)
		tR = append(tR, v)

		permuteBfs(d, tR, tmp)
	}

}

func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{
			nums,
		}
	}

	d := &PermuteData{data: make([][]int, 0)}
	tmp := make([]int, 0)
	permuteBfs(d, tmp, nums)
	return d.data
}
