package main

func partition(nums []int, start int, end int) int {
	pivot := nums[end]
	i := start - 1
	for j := start; j < end; j++ {
		if nums[j] < pivot {
			i += 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[end], nums[i+1] = nums[i+1], nums[end]

	return i + 1
}

func quick_sort(nums []int, start int, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quick_sort(nums, start, pivot-1)
		quick_sort(nums, pivot+1, end)
	}
}

func findContentChildren(g []int, s []int) int {
	quick_sort(g, 0, len(g)-1)
	quick_sort(s, 0, len(s)-1)

	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
			j++
		} else {
			j++
		}
	}
	return i
}

func main() {

}
