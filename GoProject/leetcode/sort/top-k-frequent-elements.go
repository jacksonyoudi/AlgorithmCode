package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	if len(nums) <= k {
		return nums
	}

	l := 0
	for _, num := range nums {
		i, ok := m[num]
		if ok {
			m[num] = i + 1
			if i+1 > l {
				l = i + 1
			}
		} else {
			m[num] = 1
			if l == 0 {
				l = 1
			}
		}
	}

	ints := make([]int, l+1)

	fmt.Println(m)
	for index, cnt := range m {
		ints[cnt] = index
	}

	fmt.Println(ints)
	i := l
	j := 0
	result := []int{}

	for i >= 0 && j < k {
		fmt.Println(i)
		if ints[i] > 0 {
			result = append(result, ints[i])
			j++
		}
		i--
	}

	return result
}

func tmp() {
	a := []int{1, 2, 3, 4, 5}

	for i, _ := range a {
		tmp := make([]int, 0)
		tmp = append(tmp,a[0:i]...)
		tmp = append(tmp,a[i+1:]...)
		fmt.Println(tmp)
	}

}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	tmp()
}
