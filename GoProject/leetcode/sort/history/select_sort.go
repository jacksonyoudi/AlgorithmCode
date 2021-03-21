package main

func selectionSort(nums []int) {
	l := len(nums)
	var minIndex int
	for i := 0; i < l-1; i++ {
		minIndex = 1
		for j := i; j < l; j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}

}

func main() {
	a := []int{5, 2, 3, 1}
	bubbleSort(a)
	for _, v := range a {
		println(v)
	}
}
