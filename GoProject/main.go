package main

func helper(nums []int, k int) bool {
	if findGCD(nums) == k {
		return true
	} else {
		return false
	}

}

func findGCD(nums []int) int {
	max := 1
	min := int(1e3)
	for _, n := range nums {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return gcd(max, min)
}

func gcd(a, b int) int {
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}
	return a
}

func subarrayGCD(nums []int, k int) int {
	cnt := 0
	l := len(nums)

	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {

			println("i, j, {} {} ", i, j)
			if ok := helper(nums[i:j+1], k); ok {
				println("ok,{}, {}", i, j)
				cnt += 1
			}
		}
	}

	return cnt

}

func main() {
	ints := []int{12, 9, 6}
	println(findGCD(ints))

}
