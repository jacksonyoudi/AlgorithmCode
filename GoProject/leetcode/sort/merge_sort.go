package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, 0)
	middle := len(arr) / 2
	left := mergeSort(arr[0:middle])
	right := mergeSort(arr[middle:len(arr)])

	i, j := 0, 0
	ll, lr := len(left), len(right)

	for i < ll && j < lr {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i += 1
		} else {
			result = append(result, right[j])
			j += 1
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	fmt.Println(result)
	return result
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	result := mergeSort(arr)
	fmt.Println(result)
}
