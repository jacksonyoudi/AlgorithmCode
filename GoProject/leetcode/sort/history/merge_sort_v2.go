package main

import "fmt"

func merge_sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	middle := len(arr) / 2
	left := merge_sort(arr[:middle])
	right := merge_sort(arr[middle:])

	l, r := 0, 0
	ll, lr := len(left), len(right)

	res := make([]int, 0)

	for l < ll && r < lr {
		if left[l] < right[r] {
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
		}
	}

	res = append(res, left[l:]...)
	res = append(res, right[r:]...)

	return res
}

func main() {

	arr := []int{10, 7, 8, 9, 1, 5}
	ints := merge_sort(arr)

	fmt.Println(ints)
}
