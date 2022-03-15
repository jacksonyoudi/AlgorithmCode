package array

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	all := m + n

	mid := all / 2
	i, j := 0, 0

	cnt := 0

	midValue := -1

	for i < m && j < n {

		if nums1[i] < nums2[j] {
			cnt++
			i++
			if cnt >= mid {
				midValue = nums1[i]
				break
			}
		} else {
			j++
			cnt++
			if cnt >= mid {
				midValue = nums2[j]
				break
			}
		}
	}

	return float64(midValue)


}
