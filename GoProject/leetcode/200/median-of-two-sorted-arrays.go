package _200

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	l1, l2 := len(nums1), len(nums2)
	if l1+l2 == 0 {
		return 0.0
	}

	result := make([]int, 0)

	i, j := 0, 0
	for i < l1 && j < l2 {
		if nums1[i] <= nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}

	result = append(result, nums1[i:]...)
	result = append(result, nums2[j:]...)

	ll := len(result)
	if ll%2 == 0 {
		return float64(result[ll/2]+result[ll/2-1]) / 2.0
	} else {
		return float64(result[ll/2])
	}

}
