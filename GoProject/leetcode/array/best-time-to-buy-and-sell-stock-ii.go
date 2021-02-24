package array

func maxProfit(prices []int) int {
	slow, quick := 0, 0
	sum := 0
	l := len(prices)
	for i := 0; i < l-1; i++ {
		if prices[i] <= prices[i+1] {
			quick += 1
		} else {
			sum += (prices[quick] - prices[slow])
			slow = i +1
			quick = slow
		}

	}
	sum += (prices[quick] - prices[slow])

	return sum

}
