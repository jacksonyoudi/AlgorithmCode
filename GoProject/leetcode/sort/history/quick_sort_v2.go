package main

func partitions(arr []int, start int, end int) int {
	i := start - 1 // 最小索引
	pivot := arr[end]

	for j := start; j < end; j++ {
		if arr[j] < pivot {
			i += 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[end] = arr[end], arr[i+1]
	return i + 1

}

func quick_sort(arr []int, start int, end int) {
	if start < end {
		pi := partitions(arr, start, end)
		quick_sort(arr, start, pi-1)
		quick_sort(arr, pi+1, end)

	}

}

func main() {

	arr := []int{10, 7, 8, 9, 1, 5}
	quick_sort(arr, 0, len(arr)-1)

	for _, v := range arr {
		println(v)
	}
}
