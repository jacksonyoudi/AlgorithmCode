package main

import (
	"fmt"
)

func candy(ratings []int) int {
	childs := len(ratings)
	candies := make([]int, childs)
	for i := 0; i < childs; i++ {
		candies[i] = 1
	}
	for i := 1; i < childs; i++ {
		if ratings[i-1] < ratings[i] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := childs - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			if candies[i-1] <= candies[i] {
				candies[i-1] = candies[i] + 1
			}
		}
	}
	sum := 0

	for _, i := range candies {
		sum += i
	}

	return sum
}

func main() {
	c := []int{1,2,87,87,87,2,1}
	fmt.Println(candy(c))
}
