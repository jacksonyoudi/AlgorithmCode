package main

import "fmt"

func partition(arr []int, start int, end int) int {
	i := start - 1
	pivot := arr[end]

	for j := start; j < end; j++ {
		if arr[j] <= pivot {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[end], arr[i+1] = arr[i+1], arr[end]
	return i + 1

}

func quickSort(arr []int, start int, end int) {
	if start < end {
		pivot := partition(arr, start, end)

		quickSort(arr, start, pivot-1)
		quickSort(arr, pivot+1, end)
	}
}
func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
