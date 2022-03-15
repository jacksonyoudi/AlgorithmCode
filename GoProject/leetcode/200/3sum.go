package _200

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) <= 2 {
		return [][]int{}
	}
	sort.Ints(nums)

	lm := len(nums) - 1
	l1, l2 := 0, len(nums)-1
	if nums[l1] >= 0 || nums[l2] <= 0 {
		return [][]int{}
	}

	result := make([][]int, 0)

	for l1 < l2 && l1 < lm && l2 > 0 {
		if nums[l1]+nums[l2] >= 0 {

			for l := l1 + 1; l < l2; l++ {
				if nums[l] > 0 {
					l2--
					break
				}

				if nums[l1]+nums[l2]+nums[l] == 0 {
					result = append(result, []int{l1, l, l2})
					if nums[l] > 0 {
						l2--
					} else {
						l1++
					}
					break
				}
			}
		} else if nums[l1]+nums[l2] <= 0 {
			for l := l2 - 1; l < l2; l-- {
				if nums[l] < 0 {
					l1++
				}


				if nums[l1]+nums[l2]+nums[l] == 0 {
					result = append(result, []int{l1, l, l2})
					if nums[l] > 0 {
						l2--
					} else {
						l1++
					}
					break
				}

			}
		}

	}
	return result

}
