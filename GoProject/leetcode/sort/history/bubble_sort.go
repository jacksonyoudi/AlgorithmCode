package main

func bubbleSort(nums []int) {
	l := len(nums)

	swap := true

	for i := 0; i < l-1; i++ {
		if !swap {
			break
		}
		swap = false
		for j := 0; j < l-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swap = true
			}
		}
	}
}

func main() {
	a := []int{5, 2, 3, 1}
	bubbleSort(a)
	for _, v := range a {
		println(v)
	}
}
