package difference

//
//bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
// [10,55,45,25,25]
func corpFlightBookings(bookings [][]int, n int) []int {
	//  初始化一个全是0的数组
	nums := make([]int, n)

	// 构造等差数组
	for _, booking := range bookings {
		nums[booking[0]-1] += booking[2]
		if booking[1] < n {
			nums[booking[1]] -= booking[2]
		}
	}


	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	return nums

}
